#!/usr/bin/env python3

# Copyright 2019 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""
Output devstats repo_groups.sql based on subproject defintions in sigs.yaml

This is likely missing a few repos because:
    - some repos lack an owner (eg: kubernetes/kubernetes)
    - it doesn't enumerate all repos from all kubernetes-owned orgs
    - it ignores the fact that committees can own repos, only grouping by sig

The sql generated is NOT intended to overwrite/replace the file that lives at
github.com/cncf/devstats/scripts/kubernetes/repo_groups.sql, but instead aid a
human in doing some manual updates to the file. Future improvements to this
script could eliminate that part of the process, but it's where we are today.
"""

import argparse
import ruamel.yaml as yaml
import json
import re
import sys

update_gha_repos_template = """
update gha_repos set repo_group = 'SIG {}' where name in (
{}
);
"""

def repos_from_sig(sig):
    """Returns a list of org/repos given a sig"""
    repos = {}
    subprojects = sig.get('subprojects', [])
    if subprojects is None:
        subprojects = []
    for sp in subprojects:
        for uri in sp['owners']:
            owners_path = re.sub(r"https://raw.githubusercontent.com/(.*)/master/(.*)",r"\1/\2",uri)
            path_parts = owners_path.split('/')
            # org/repo is owned by sig if org/repo/OWNERS os in one of their subprojects
            if path_parts[2] == 'OWNERS':
                repo = '/'.join(path_parts[0:2])
                repos[repo] = True
    return sorted(repos.keys())

def write_repo_groups_sql(sigs, fp):
    for sig in sigs['sigs']:
        repos = repos_from_sig(sig)
        if len(repos):
            fp.write(
                update_gha_repos_template.format(
                    sig['name'],
                    ',\n'.join(['  \'{}\''.format(r) for r in repos])))

def main(sigs_yaml, repo_groups_sql):
    with open(sigs_yaml) as fp:
        sigs = yaml.round_trip_load(fp)

    if repo_groups_sql is not None:
        with open(repo_groups_sql, 'w') as fp:
            write_repo_groups_sql(sigs, fp)
    else:
        write_repo_groups_sql(sigs, sys.stdout)

if __name__ == '__main__':
    PARSER = argparse.ArgumentParser(
        description='Do things with sigs.yaml')
    PARSER.add_argument(
        '--sigs-yaml',
        default='./sigs.yaml',
        help='Path to sigs.yaml')
    PARSER.add_argument(
        '--repo-groups-sql',
        help='Path to output repo_groups.sql if provided')
    ARGS = PARSER.parse_args()

    main(ARGS.sigs_yaml, ARGS.repo_groups_sql)

