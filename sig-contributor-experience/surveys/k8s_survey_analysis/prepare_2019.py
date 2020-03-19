import pandas as pd
import numpy as np

fn = '2019_survey/2019 Kubernetes Contributor Experience Survey PUBLIC.csv'

contribute_header = "What areas of Kubernetes do you contribute to? Please check all that apply."
blockers_header = "Please rate any challenges to the listed steps of the contribution process"
agree_header = "Do you agree with the following statements (1 - strongly disagree, 5 - strongly agree):"
attend_header = "Which of the below would make you likely to attend more of the Community Meetings? Check all that apply."
most_important_proj_header = "Some of the major projects SIG Contributor Experience is working on are listed below, rank the ones that are most important to you (and/or your SIG)"
use_freq_header = "Of our various communications channels, please rate which ones you use and/or check most frequently on a 1-5 scale, where 1 is “never”, 3 is “several times a month” and 5 is “every day”."
news_header = "Which of these channels is most likely to reach you first for news about decisions, changes, additions, and/or announcements to the contributor process or community matters?"

def map_blocker_and_usefreq_vals(val):
    try:
        return int(val)
    except ValueError:
        return int(val[0])

def process_header(df):
    columns = list(df.columns)
    new_columns = [None]*len(columns)
    for i, col in enumerate(columns):
        if col[1].startswith("Unnamed") or col[1] == "Response":
            new_columns[i] = col[0]
            continue

        # Find the starting column for the multilabel responses (checkboxes)
        # that were also in the 2018 survey
        if col[0] == blockers_header:
            blockers_i = i
        elif col[0] == contribute_header:
            contribute_i = i
        elif col[0] == news_header:
            news_i = i
        elif col[0] == use_freq_header:
            use_freq_i = i
        elif col[0] == most_important_proj_header:
            most_important_proj_i = i
        elif col[0] == agree_header: # Starting columns for multilabel responses that weren't in the 2018 survey.
            agree_i = i
        elif col[0] == attend_header:
            attend_i = i
        #elif col[0] == unattendance_header:
        #    unattendance_i = i
        else: # Handle open ended responses
            new_columns[i] = col[0]

    def prefix_cols(header, header_i, prefix):
        i = header_i
        while i < len(columns) and (columns[i][0].startswith("Unnamed") or columns[i][0] == header):
            new_columns[i] = "{} {}".format(prefix, columns[i][1])
            i += 1

    prefix_cols(contribute_header, contribute_i, "Contribute:")
    prefix_cols(blockers_header, blockers_i, "Blocker:")
    prefix_cols(news_header, news_i, "Check for news:")
    prefix_cols(use_freq_header, use_freq_i, "Use freq:")
    prefix_cols(most_important_proj_header, most_important_proj_i, "Most Important Project:")

    prefix_cols(agree_header, agree_i, "Agree:")
    prefix_cols(attend_header, attend_i, "Would attend if:")

    df.columns = new_columns

def get_df(file_name=None):
    fn = '2019_survey/2019 Kubernetes Contributor Experience Survey PUBLIC.csv'
    if file_name:
        fn = file_name   

    df = pd.read_csv(fn, header=[0,1], skipinitialspace=True)
    process_header(df)

    df = df.rename(columns={
        "How long have you been contributing to Kubernetes?": "Contributing_Length",
        "What level of the Contributor Ladder do you consider yourself to be on? (pick the highest if you are in multiple OWNERs files)": "Level_of_Contributor",
        "What level of the Contributor Ladder do you consider yourself to be on?  (pick the highest if you are in multiple OWNERs files)": "Level_of_Contributor",
        "What region of the world are you in?": "World_Region",
        "Are you interested in advancing to the next level of the Contributor Ladder?": "Interested_in_next_level",
        "How many other open source projects not in the Kubernetes ecosystem do you contribute to? (example: nodejs, debian)":"Contribute_to_other_OSS",
        "Does your employer support your contributions to Kubernetes?":"Upstream_supported_at_employer",
        "Blocker: Other (please specify)": "Other blockers (please specify)",
        "What region of the world are you in?": "World Region",
    })

    def map_blocker_and_usefreq_vals(val):
        try:
            return int(val)
        except ValueError:
            return int(val[0])

    #Clean Data
    for x in df.columns:
        if x.startswith("Useful:"):
            df = df.assign(**{x: df[x].fillna(0)})
        if x.startswith("Contribute:") or x.startswith("Check for news:") or x.startswith("Attended:") or x.startswith("Attending:") or x.startswith("Would attend if:"):
            df = df.assign(**{x: np.where(df[x].isna(),0,1)})
        if x.startswith('Upstream'):
            df = df.assign(**{x: df[x].fillna("Didn't Answer")})
        if x.startswith("Blocker:") and x != "Blocker: Other (please specify)":
            df[x] = df[x].map(map_blocker_and_usefreq_vals, na_action="ignore")
        if x.startswith("Use freq:") or x.startswith("Agree:"):
            df[x] = df[x].map(map_blocker_and_usefreq_vals, na_action="ignore")
        

    df = df.rename(columns= {x:x.replace(" ","_").replace("?", "").replace('Most_Important_Project','Most_Important_Proj').replace('Most_Important_Prj','Most_Important_Proj') for x in df.columns})

    x = pd.to_datetime(df.End_Date)
    df = df.assign(date_taken = x.dt.date)

    return df

# TODO NOTE I should only be dropping these at plot time
#df.dropna(subset=["Level_of_Contributor",
#                  "Interested_in_next_level",
#                  "Upstream_supported_at_employer"], inplace=True)

