import pandas as pd
import numpy as np


convert_2018_to_2019 = {
    'Blocker:_Code/Doc_review':'Blocker:_Code/Documentation_review',
    'Blocker:_GH_tools&processes_(not_our_customized_tooling)': 'Blocker:_GitHub_tools_and_processes_(not_our_customized_tooling)',
    'Blocker:_Finding_a/the_right_SIG': 'Blocker:_Finding_the_right_SIG_for_your_contributions',
    'Blocker:_Finding_issues_to_work_on': 'Blocker:_Finding_appropriate_issues_to_work_on',
    'Blocker:_Setting_up_dev_env': 'Blocker:_Setting_up_development_environment',
    'Use_freq:_Zoom_Mtgs': 'Use_freq:_Zoom_video_conferencing/meetings',
    'Use_freq:_GH_(comments,_issues,_prs)': 'Use_freq:_Discussions_on_Github_Issues_and_PRs',
    'Use_freq:_Unofficial(Twitter,_Reddit,_etc.)':'Use_freq:_Unofficial_channels_(IRC,_WeChat,_etc.)',
    'Use_freq:_YT_Recordings': 'Use_freq:_YouTube_recordings_(community_meetings,_SIG/WG_meetings,_etc.)',
    'Use_freq:_GDocs/Forms/Sheets,_etc_(meeting_agendas,_etc)': 'Use_freq:_Google_Docs/Forms/Sheets,_etc_(meeting_agendas,_etc)',
    'Contribute:_code_to_k/k': 'Contribute:_Core_code_inside_of_kubernetes/kubernetes',
    'Contribute:_code_in_a_k/*_GH_org': 'Contribute:_Code_inside_of_another_repo_in_the_Kubernetes_GitHub_Org_(example:_/kubernetes-sigs,_kubernetes/website,_etc)',
    'Contribute:_Docs':'Contribute:_Documentation',
    'Contribute:_Testing_and_CI':'Contribute:_Testing_&_Infrastructure',
    'Contribute:_Related_projects_(Kubeadm,_Helm,_container_runtimes,_etc.)': 'Contribute:_Related_projects_(Helm,_container_runtimes,_other_CNCF_projects,_etc.)',
    'Contribute:_Not_yet': 'Contribute:_Don’t_contribute_yet,_hoping_to_start_soon',
    'Contribute:_Other': 'Contribute:_Other_(please_specify)',
    'Level_of_Contributor_Laddor':'Level_of_Contributor',
    'Most_Important_Proj:_Mentoring_programs':'Most_Important_Proj:_Mentoring_programs_for_all_contributor_levels/roles\xa0(https://git.k8s.io/community/community-membership.md)',
    'Most_Important_Proj:_GH_Mgmt':'Most_Important_Proj:_GitHub_Management',
    'Most_Important_Proj:_Contributor_Summits':'Most_Important_Proj:_Delivering_valuable_contributor_summits_at_relevant_events',
    'Most_Important_Proj:_Keeping_community_safe': 'Most_Important_Proj:_Keeping_our_community_safe_on_our_various_communication_platforms_through_moderation_guidelines_and_new_approaches',
    'Check_for_news:_k-dev_ML':'Check_for_news:_kubernetes-dev@_mailing_list',
    'Check_for_news:_discuss.kubernetes.io':'Check_for_news:_Dedicated_discuss.k8s.io_forum_for_contributors',
    'Check_for_news:_contribex_ML':'Check_for_news:_kubernetes-sig-contribex@\xa0mailing_list',
    'Check_for_news:_Slack':'Check_for_news:_#kubernetes-dev,_#sig-foo,_#sig-contribex_slack',
    'Check_for_news:_Twitter_read_first_':'Check_for_news:_@kubernetesio_twitter',
    'Check_for_news:_Kubernetes_blog_read_first_':'Check_for_news:_Kubernetes_blog',
    'Check_for_news:_k/community_repo_in_GH_(Issues_and/or_PRs)_read_first':'Check_for_news:_kubernetes/community_repo_in_GitHub_(Issues_and/or_PRs)',
    'Check_for_news:_Other':'Check_for_news:_Other_(please_specify)',
    'Attended:_#_of_ContribSummits':'How_many_Kubernetes_Contributor_Summits_have_you_attended',
    'HelpWanted_&/or_GoodFirstIssue_label_usage':'Do_you_use_the\xa0Help_Wanted_and/or_Good_First_Issue_labels_on_issues_you_file_to_find_contributors',
    'Watched_or_participated_in_MoC':'Have_you_watched_or_participated_in_an_episode_of_our_YouTube_mentoring_series_Meet_Our_Contributors_If_you_have_specific_suggestions,_leave_them_at_the_end_of_the_survey.',
    'Make_project_easier_to_contribute':'Are_there_specific_ways_the_project_could_make_contributing_easier_for_you'
}

contrib_length_2018_to_2019 = {
    '1-2 years': 'one to two years',
    '2-3 years': 'two to three years',
    '3+ years': 'three+ years',
    '6 months-1 year':'less than one year',
    'Just started': 'less than one year'
}

ladder_level_2018_to_2019 = {
    "Approver": "approver",
    "Had no idea this was even a thing": "there's a contributor ladder?",
    "Org Member": "member",
    "Reviewer": "reviewer",
    "I’m not an org member yet, but working on it": "not yet a member but working on it",
    "Subproject Owner": "subproject owner"
}

employer_2018_to_2019 = {
    "It’s complicated": "it's complicated.",
    "It’s entirely on my own time": "no, I need to use my own time",
    "Yes, it’s part of my job": "yes, I can contribute on company time",
    'No, but I’m able to use “free” time at work': "yes, I can contribute on company time"
}

oss_projects_2018_to_2019 = {
    'None, Kubernetes is my first one!': 'this is my first open source project!',
    'One more':'1 other',
    '2-4' : '2 or more',
    '4+': '2 or more' 
}

help_wanted_2018_to_2019 = {
    "No, because I didn't know they were there": "No",
    "No, because I don't think my issues qualify": "No",
    'Not as much as I should because I forget' : "Rarely (for reasons)" 
}

next_level_interest_2018_2019 = {
    'Yes, but would like mentorship.': 'if I had help/mentoring/support',
    'Yes, but not sure I have time.': 'if I had more free time',
    'Yes, doing it on my own.': 'yes',
    "No, I'm already an owner": 'no, already a subproject owner (highest level on the ladder)',
    'Not really': 'no'
}

def get_df(path):

    survey_data = pd.read_csv(path)

    #Clean Data
    for x in survey_data.columns:
        if x.startswith("Useful:"):
            survey_data = survey_data.assign(**{x: survey_data[x].fillna(0)})
        if x.startswith("Contribute:") or x.startswith("Check for news:") or x.startswith("Attended:") or x.startswith("Attending:") or x.startswith("Most Important Pr"):
            survey_data = survey_data.assign(**{x: np.where(survey_data[x].isna(),0,1)})
        if x.startswith('Upstream'):
            survey_data = survey_data.assign(**{x: survey_data[x].fillna("Didn't Answer")})
    
   

    survey_data = survey_data.rename(columns= {x:x.replace(" ","_").replace("?", "").replace('Most_Important_Project','Most_Important_Proj').replace('Most_Important_Prj','Most_Important_Proj') for x in survey_data.columns})
    
    survey_data = survey_data.drop('Use_freq:_discuss.kubernetes.io',axis=1)

    x = pd.to_datetime(survey_data.End_Date)
    survey_data = survey_data.assign(date_taken = x.dt.date)
    survey_data = survey_data.assign(Contributing_Length = survey_data['Contributing_Length'].apply(contrib_length_2018_to_2019.get))
   
    survey_data = survey_data.rename(columns=convert_2018_to_2019)

    survey_data = survey_data.assign(Level_of_Contributor = survey_data['Level_of_Contributor'].apply(lambda x: ladder_level_2018_to_2019.get(x,x)))
    survey_data = survey_data.assign(Upstream_supported_at_employer = survey_data['Upstream_supported_at_employer'].apply(lambda x: employer_2018_to_2019.get(x,x)))
    survey_data = survey_data.assign(Interested_in_next_level = survey_data['Interested_in_next_level'].apply(lambda x: next_level_interest_2018_2019.get(x,x) ))
    survey_data = survey_data.assign(Contribute_to_other_OSS = survey_data['Contribute_to_other_OSS'].apply(lambda x: oss_projects_2018_to_2019.get(x,x)))
    survey_data.loc[:,'Do_you_use_the\xa0Help_Wanted_and/or_Good_First_Issue_labels_on_issues_you_file_to_find_contributors'] = survey_data['Do_you_use_the\xa0Help_Wanted_and/or_Good_First_Issue_labels_on_issues_you_file_to_find_contributors'].apply(lambda x: help_wanted_2018_to_2019.get(x,x))

    return survey_data
