```python
import pandas as pd
import plotnine as p9
from textwrap import wrap
import sys
sys.path.append("../")

from k8s_survey_analysis.plot_utils import (
    make_likert_chart_multi_year,
    make_bar_chart_multi_year
)

from k8s_survey_analysis import prepare_2019, prepare_2018
```


```python
pd.options.display.max_columns = 999
pd.options.display.max_rows = 999
```


```python
# Match Field Names and Multiple Choice Options as much as possible
df_2019 = prepare_2019.get_df('../2019/contribex-survey-2019.csv').assign(year=2019)
df_2018 = prepare_2018.get_df('../2018/contribex-survey-2018.csv').assign(year=2018)
```

## Build a single dataset for comparison

In this notebook we will compare responses between the 2018 and 2019 surveys. Only those questions which appeared in both surveys are useful to compare. Some questions that appeared in both surveys were different enough in their format to make comparisons not useful.


```python
shared_columns = set(df_2019.columns).intersection(df_2018.columns)
survey_data = pd.concat([df_2019[shared_columns], df_2018[shared_columns]])
```

## Compare univariate data between 2018 and 2019


```python
length_year_totals = (
    survey_data[survey_data["Contributing_Length"].notnull()]
    .groupby(["year"])
    .count()["Collector_ID"]
    .reset_index()
)

length_counts = (
    survey_data.groupby(["Contributing_Length", "year"])
    .count()["Collector_ID"]
    .reset_index()
)

length_percents = length_counts.merge(length_year_totals, on="year")

length_percents = length_percents.assign(
    percent=length_percents["Collector_ID_x"] / length_percents["Collector_ID_y"]
)
```


```python
(
    p9.ggplot(
        length_percents,
        p9.aes(x="Contributing_Length", fill="factor(year)", y="percent"),
    )
    + p9.geom_bar(position="dodge", stat="identity")
    + p9.theme(axis_text_x=p9.element_text(angle=45))
    + p9.scale_x_discrete(
        limits=[
            "less than one year",
            "one to two years",
            "two to three years",
            "three+ years",
        ]
    )
    + p9.ggtitle("Number of Contributors by Length of Contribution")
    + p9.xlab("Length of Contribution")
    + p9.ylab("Number of Contributors")
)
```


![png](https://drive.google.com/uc?export=view&id=1jMfUFhWCxhI3VdhsFJyKH7Tq7gBCmc_C)


As expected, due to the advertisement of the survey on Twitter, a higher proportion of respondents are newer contributors this year.


```python
level_year_totals = survey_data[survey_data['Level_of_Contributor'].notnull()].groupby(['year']).count()['Collector_ID'].reset_index()
level_counts = survey_data.groupby(['Level_of_Contributor','year']).count()['Collector_ID'].reset_index()
level_percents = level_counts.merge(level_year_totals,on='year')
level_percents = level_percents.assign(percent = level_percents['Collector_ID_x']/level_percents['Collector_ID_y'])


```


```python
(
    p9.ggplot(
        level_percents,
        p9.aes(x="Level_of_Contributor", fill="factor(year)", y="percent"),
    )
    + p9.geom_bar(position="dodge", stat="identity")
    + p9.theme(axis_text_x=p9.element_text(angle=45, ha="right"))
    + p9.ggtitle("Number of Contributors by Contributor Level")
    + p9.xlab("Contributor Level")
    + p9.ylab("Proportion of Contributors")
)
```


![png](https://drive.google.com/uc?export=view&id=1MsS393t45gOKqvUp_IrmZzOe58MHB6aa)



A larger proportion of respondents this year are in higher roles in the contributor ladder. It appears most of these are due to the smaller proportion of respondents who are not a member for one reason or another.


```python
oss_counts = (
    survey_data.groupby(["year", "Contribute_to_other_OSS"])
    .count()["Collector_ID"]
    .reset_index()
)
oss_year_totals = (
    survey_data[survey_data["Contribute_to_other_OSS"].notnull()]
    .groupby(["year"])
    .count()["Collector_ID"]
    .reset_index()
)

oss_percents = oss_counts.merge(oss_year_totals, on="year")
oss_percents = oss_percents.assign(
    percent=oss_percents["Collector_ID_x"] / oss_percents["Collector_ID_y"]
)
```


```python
(
    p9.ggplot(
        oss_percents,
        p9.aes(x="Contribute_to_other_OSS", fill="factor(year)", y="percent"),
    )
    + p9.geom_bar(position="dodge", stat="identity")
    + p9.theme(axis_text_x=p9.element_text(angle=45, ha="right"))
    + p9.scale_x_discrete(
        limits=["this is my first open source project!", "1 other", "2 or more"]
    )
    + p9.ggtitle("Participation in Other Open Source Projects")
    + p9.xlab("Number of other OSS Projects")
    + p9.ylab("Number of Contributors")
)
```


![png](https://drive.google.com/uc?export=view&id=1MsS393t45gOKqvUp_IrmZzOe58MHB6aa)


There is a small increase in the proportion of contributors for whom Kuberetes is their first open source project. Because the change is so small, no major changes should be based on this finding.


```python
counts = survey_data.groupby(["Upstream_supported_at_employer", "year"]).count()[
    "Respondent_ID"
]

total = survey_data.groupby(["year"]).count()["Respondent_ID"]

employee_percents = (
    counts.to_frame().reset_index().merge(total.to_frame().reset_index(), on="year")
)

employee_percents = employee_percents.assign(
    percent=employee_percents["Respondent_ID_x"] / employee_percents["Respondent_ID_y"]
)

(
    p9.ggplot(
        employee_percents,
        p9.aes(x="Upstream_supported_at_employer", fill="factor(year)", y="percent"),
    )
    + p9.geom_bar(position=p9.position_dodge(preserve="single"), stat="identity")
    + p9.labs(
        title="Support at Employer", fill="Year", y="Proportion", x="Support Level"
    )
    + p9.theme(axis_text_x=p9.element_text(angle=45, ha="right"))
)
```


![png](https://drive.google.com/uc?export=view&id=1zyRmPhyGMLUJR2oisgRiqzMDHhB_z8mi)




This question was a required question in 2019 but not in 2018. In addition, the student option was not present in this year's survey. Nonetheless, the proportion of contributors supported by their employer dropped slightly this year. It is difficult to make any conclusions due to the change in the question, but this is a trend to keep a close eye on in future surveys.

## Cross-tabulations


```python
pd.crosstab(
    survey_data.World_Region, [survey_data.Level_of_Contributor, survey_data.year]
)
```




<div>
<table border="1" class="dataframe">
  <thead>
    <tr>
      <th>Level_of_Contributor</th>
      <th colspan="2" halign="left">approver</th>
      <th colspan="2" halign="left">member</th>
      <th colspan="2" halign="left">not yet a member but working on it</th>
      <th colspan="2" halign="left">reviewer</th>
      <th colspan="2" halign="left">subproject owner</th>
      <th colspan="2" halign="left">there's a contributor ladder?</th>
    </tr>
    <tr>
      <th>year</th>
      <th>2018</th>
      <th>2019</th>
      <th>2018</th>
      <th>2019</th>
      <th>2018</th>
      <th>2019</th>
      <th>2018</th>
      <th>2019</th>
      <th>2018</th>
      <th>2019</th>
      <th>2018</th>
      <th>2019</th>
    </tr>
    <tr>
      <th>World_Region</th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>Africa</th>
      <td>0</td>
      <td>0</td>
      <td>0</td>
      <td>2</td>
      <td>0</td>
      <td>2</td>
      <td>0</td>
      <td>0</td>
      <td>0</td>
      <td>0</td>
      <td>0</td>
      <td>0</td>
    </tr>
    <tr>
      <th>Asia</th>
      <td>2</td>
      <td>2</td>
      <td>4</td>
      <td>5</td>
      <td>8</td>
      <td>15</td>
      <td>3</td>
      <td>5</td>
      <td>2</td>
      <td>0</td>
      <td>1</td>
      <td>3</td>
    </tr>
    <tr>
      <th>Europe</th>
      <td>5</td>
      <td>8</td>
      <td>11</td>
      <td>15</td>
      <td>21</td>
      <td>23</td>
      <td>2</td>
      <td>9</td>
      <td>2</td>
      <td>6</td>
      <td>4</td>
      <td>10</td>
    </tr>
    <tr>
      <th>North America</th>
      <td>10</td>
      <td>12</td>
      <td>19</td>
      <td>22</td>
      <td>26</td>
      <td>27</td>
      <td>8</td>
      <td>5</td>
      <td>16</td>
      <td>22</td>
      <td>14</td>
      <td>11</td>
    </tr>
    <tr>
      <th>Oceania</th>
      <td>0</td>
      <td>0</td>
      <td>0</td>
      <td>1</td>
      <td>0</td>
      <td>1</td>
      <td>0</td>
      <td>1</td>
      <td>0</td>
      <td>0</td>
      <td>0</td>
      <td>1</td>
    </tr>
    <tr>
      <th>South America</th>
      <td>0</td>
      <td>0</td>
      <td>0</td>
      <td>1</td>
      <td>0</td>
      <td>0</td>
      <td>0</td>
      <td>0</td>
      <td>1</td>
      <td>1</td>
      <td>1</td>
      <td>0</td>
    </tr>
  </tbody>
</table>
</div>




```python
pd.crosstab(
    survey_data.Contributing_Length,
    [survey_data.Level_of_Contributor, survey_data.year],
).loc[["less than one year", "one to two years", "two to three years", "three+ years"]]
```




<div>
<table border="1" class="dataframe">
  <thead>
    <tr>
      <th>Level_of_Contributor</th>
      <th colspan="2" halign="left">approver</th>
      <th colspan="2" halign="left">member</th>
      <th colspan="2" halign="left">not yet a member but working on it</th>
      <th colspan="2" halign="left">reviewer</th>
      <th colspan="2" halign="left">subproject owner</th>
      <th colspan="2" halign="left">there's a contributor ladder?</th>
    </tr>
    <tr>
      <th>year</th>
      <th>2018</th>
      <th>2019</th>
      <th>2018</th>
      <th>2019</th>
      <th>2018</th>
      <th>2019</th>
      <th>2018</th>
      <th>2019</th>
      <th>2018</th>
      <th>2019</th>
      <th>2018</th>
      <th>2019</th>
    </tr>
    <tr>
      <th>Contributing_Length</th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>less than one year</th>
      <td>0</td>
      <td>3</td>
      <td>1</td>
      <td>15</td>
      <td>24</td>
      <td>42</td>
      <td>1</td>
      <td>4</td>
      <td>0</td>
      <td>0</td>
      <td>9</td>
      <td>18</td>
    </tr>
    <tr>
      <th>one to two years</th>
      <td>2</td>
      <td>8</td>
      <td>10</td>
      <td>19</td>
      <td>1</td>
      <td>18</td>
      <td>5</td>
      <td>7</td>
      <td>5</td>
      <td>5</td>
      <td>5</td>
      <td>4</td>
    </tr>
    <tr>
      <th>two to three years</th>
      <td>3</td>
      <td>6</td>
      <td>6</td>
      <td>5</td>
      <td>4</td>
      <td>6</td>
      <td>2</td>
      <td>4</td>
      <td>7</td>
      <td>8</td>
      <td>0</td>
      <td>2</td>
    </tr>
    <tr>
      <th>three+ years</th>
      <td>9</td>
      <td>5</td>
      <td>0</td>
      <td>7</td>
      <td>0</td>
      <td>2</td>
      <td>0</td>
      <td>5</td>
      <td>9</td>
      <td>16</td>
      <td>1</td>
      <td>1</td>
    </tr>
  </tbody>
</table>
</div>




```python
pd.crosstab(
    survey_data.Contributing_Length,
    [survey_data.Contribute_to_other_OSS, survey_data.year],
).loc[
    ["less than one year", "one to two years", "two to three years", "three+ years"],
    ["this is my first open source project!", "1 other", "2 or more"],
]
```




<div>
<table border="1" class="dataframe">
  <thead>
    <tr>
      <th>Contribute_to_other_OSS</th>
      <th colspan="2" halign="left">1 other</th>
      <th colspan="2" halign="left">2 or more</th>
      <th colspan="2" halign="left">this is my first open source project!</th>
    </tr>
    <tr>
      <th>year</th>
      <th>2018</th>
      <th>2019</th>
      <th>2018</th>
      <th>2019</th>
      <th>2018</th>
      <th>2019</th>
    </tr>
    <tr>
      <th>Contributing_Length</th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>less than one year</th>
      <td>9</td>
      <td>16</td>
      <td>12</td>
      <td>39</td>
      <td>14</td>
      <td>27</td>
    </tr>
    <tr>
      <th>one to two years</th>
      <td>7</td>
      <td>13</td>
      <td>15</td>
      <td>33</td>
      <td>6</td>
      <td>15</td>
    </tr>
    <tr>
      <th>two to three years</th>
      <td>5</td>
      <td>6</td>
      <td>15</td>
      <td>17</td>
      <td>2</td>
      <td>8</td>
    </tr>
    <tr>
      <th>three+ years</th>
      <td>3</td>
      <td>8</td>
      <td>14</td>
      <td>24</td>
      <td>2</td>
      <td>4</td>
    </tr>
  </tbody>
</table>
</div>




```python
pd.crosstab(
    survey_data.Level_of_Contributor,
    [survey_data.Upstream_supported_at_employer, survey_data.year],
)
```




<div>
<table border="1" class="dataframe">
  <thead>
    <tr>
      <th>Upstream_supported_at_employer</th>
      <th colspan="2" halign="left">Didn't Answer</th>
      <th>I’m a student</th>
      <th colspan="2" halign="left">it's complicated.</th>
      <th colspan="2" halign="left">no, I need to use my own time</th>
      <th colspan="2" halign="left">yes, I can contribute on company time</th>
    </tr>
    <tr>
      <th>year</th>
      <th>2018</th>
      <th>2019</th>
      <th>2018</th>
      <th>2018</th>
      <th>2019</th>
      <th>2018</th>
      <th>2019</th>
      <th>2018</th>
      <th>2019</th>
    </tr>
    <tr>
      <th>Level_of_Contributor</th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>approver</th>
      <td>3</td>
      <td>0</td>
      <td>0</td>
      <td>1</td>
      <td>5</td>
      <td>0</td>
      <td>3</td>
      <td>13</td>
      <td>14</td>
    </tr>
    <tr>
      <th>member</th>
      <td>9</td>
      <td>0</td>
      <td>1</td>
      <td>0</td>
      <td>11</td>
      <td>4</td>
      <td>13</td>
      <td>20</td>
      <td>22</td>
    </tr>
    <tr>
      <th>not yet a member but working on it</th>
      <td>22</td>
      <td>0</td>
      <td>1</td>
      <td>1</td>
      <td>10</td>
      <td>4</td>
      <td>26</td>
      <td>27</td>
      <td>32</td>
    </tr>
    <tr>
      <th>reviewer</th>
      <td>0</td>
      <td>0</td>
      <td>1</td>
      <td>0</td>
      <td>4</td>
      <td>2</td>
      <td>5</td>
      <td>10</td>
      <td>11</td>
    </tr>
    <tr>
      <th>subproject owner</th>
      <td>3</td>
      <td>0</td>
      <td>0</td>
      <td>0</td>
      <td>5</td>
      <td>0</td>
      <td>1</td>
      <td>18</td>
      <td>23</td>
    </tr>
    <tr>
      <th>there's a contributor ladder?</th>
      <td>8</td>
      <td>1</td>
      <td>0</td>
      <td>2</td>
      <td>0</td>
      <td>2</td>
      <td>10</td>
      <td>8</td>
      <td>14</td>
    </tr>
  </tbody>
</table>
</div>




```python
pd.crosstab(
    survey_data.Contributing_Length,
    [survey_data.Upstream_supported_at_employer, survey_data.year],
)
```




<div>
<table border="1" class="dataframe">
  <thead>
    <tr>
      <th>Upstream_supported_at_employer</th>
      <th colspan="2" halign="left">Didn't Answer</th>
      <th>I’m a student</th>
      <th colspan="2" halign="left">it's complicated.</th>
      <th colspan="2" halign="left">no, I need to use my own time</th>
      <th colspan="2" halign="left">yes, I can contribute on company time</th>
    </tr>
    <tr>
      <th>year</th>
      <th>2018</th>
      <th>2019</th>
      <th>2018</th>
      <th>2018</th>
      <th>2019</th>
      <th>2018</th>
      <th>2019</th>
      <th>2018</th>
      <th>2019</th>
    </tr>
    <tr>
      <th>Contributing_Length</th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>less than one year</th>
      <td>18</td>
      <td>2</td>
      <td>1</td>
      <td>1</td>
      <td>10</td>
      <td>6</td>
      <td>37</td>
      <td>9</td>
      <td>34</td>
    </tr>
    <tr>
      <th>one to two years</th>
      <td>2</td>
      <td>0</td>
      <td>2</td>
      <td>0</td>
      <td>12</td>
      <td>2</td>
      <td>16</td>
      <td>22</td>
      <td>33</td>
    </tr>
    <tr>
      <th>three+ years</th>
      <td>4</td>
      <td>0</td>
      <td>0</td>
      <td>0</td>
      <td>5</td>
      <td>0</td>
      <td>4</td>
      <td>15</td>
      <td>27</td>
    </tr>
    <tr>
      <th>two to three years</th>
      <td>6</td>
      <td>0</td>
      <td>0</td>
      <td>0</td>
      <td>8</td>
      <td>1</td>
      <td>1</td>
      <td>15</td>
      <td>22</td>
    </tr>
  </tbody>
</table>
</div>




```python
pd.crosstab(
    survey_data.Interested_in_next_level,
    [survey_data.Level_of_Contributor, survey_data.year],
)
```




<div>

<table border="1" class="dataframe">
  <thead>
    <tr>
      <th>Level_of_Contributor</th>
      <th colspan="2" halign="left">approver</th>
      <th colspan="2" halign="left">member</th>
      <th colspan="2" halign="left">not yet a member but working on it</th>
      <th colspan="2" halign="left">reviewer</th>
      <th colspan="2" halign="left">subproject owner</th>
      <th colspan="2" halign="left">there's a contributor ladder?</th>
    </tr>
    <tr>
      <th>year</th>
      <th>2018</th>
      <th>2019</th>
      <th>2018</th>
      <th>2019</th>
      <th>2018</th>
      <th>2019</th>
      <th>2018</th>
      <th>2019</th>
      <th>2018</th>
      <th>2019</th>
      <th>2018</th>
      <th>2019</th>
    </tr>
    <tr>
      <th>Interested_in_next_level</th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>if I had help/mentoring/support</th>
      <td>2</td>
      <td>1</td>
      <td>8</td>
      <td>13</td>
      <td>31</td>
      <td>18</td>
      <td>4</td>
      <td>3</td>
      <td>2</td>
      <td>0</td>
      <td>6</td>
      <td>3</td>
    </tr>
    <tr>
      <th>if I had more free time</th>
      <td>4</td>
      <td>8</td>
      <td>9</td>
      <td>8</td>
      <td>12</td>
      <td>13</td>
      <td>2</td>
      <td>3</td>
      <td>3</td>
      <td>4</td>
      <td>8</td>
      <td>6</td>
    </tr>
    <tr>
      <th>no</th>
      <td>2</td>
      <td>1</td>
      <td>3</td>
      <td>3</td>
      <td>2</td>
      <td>3</td>
      <td>2</td>
      <td>0</td>
      <td>0</td>
      <td>0</td>
      <td>4</td>
      <td>2</td>
    </tr>
    <tr>
      <th>no, already a subproject owner (highest level on the ladder)</th>
      <td>5</td>
      <td>0</td>
      <td>1</td>
      <td>0</td>
      <td>0</td>
      <td>0</td>
      <td>0</td>
      <td>0</td>
      <td>13</td>
      <td>25</td>
      <td>0</td>
      <td>0</td>
    </tr>
    <tr>
      <th>yes</th>
      <td>4</td>
      <td>11</td>
      <td>13</td>
      <td>22</td>
      <td>10</td>
      <td>34</td>
      <td>5</td>
      <td>14</td>
      <td>3</td>
      <td>0</td>
      <td>2</td>
      <td>14</td>
    </tr>
  </tbody>
</table>
</div>



## Challenges/Blockers Faced


```python
blocker_ratings = ["A frequent blocker",'Often a problem','Sometimes a problem','Rarely a problem','Not a problem']

```


```python
(
    make_likert_chart_multi_year(survey_data, "Blocker:", blocker_ratings)
    + p9.labs(
        x="Year",
        y="",
        fill="Rating",
        color="Rating",
        title="Blockers (New Contributors Included)",
    )
    + p9.theme(
        legend_margin=5,
        figure_size=(12, (9 / 4) * 3),
        strip_text_x=p9.element_text(
            margin={"t": 0.35, "b": 0.35, "l": 1, "r": 1, "units": "in"}
        ),
    )
    + p9.ylim(-1, 1)
)
```


![png](https://drive.google.com/uc?export=view&id=12kDcXR8hh8RLAcda9_mgL0-V3Jp8msP-)




```python
(
    make_likert_chart_multi_year(survey_data, "Blocker:", blocker_ratings)
    + p9.labs(
        x="Year",
        y="",
        fill="Rating",
        color="Rating",
        title="Blockers (New Contributors Included)",
    )
    + p9.theme(
        legend_margin=5,
        figure_size=(12, (9 / 4) * 3),
        strip_text_x=p9.element_text(
            margin={"t": 0.35, "b": 0.35, "l": 1, "r": 1, "units": "in"}
        ),
    )
    + p9.ylim(-1, 1)
)
```


![png](https://drive.google.com/uc?export=view&id=16C8pj5s2DN43IKe1tKe_Onn3uSDDaIct)




Respondents across the board reported feeling more blocked this past year than in 2018. The only exception is GitHub tooling. To ensure this was not due to a higher proportion of new contributors we looked at the data without them. This is show below.


```python
(
    make_likert_chart_multi_year(
        survey_data, "Blocker:", blocker_ratings, exclude_new_contributors=True
    )
    + p9.labs(
        x="Blocker",
        y="Count",
        fill="Rating",
        color="Rating",
        title="Blockers (New Contributors Excluded)",
    )
    + p9.theme(
        legend_margin=5,
        figure_size=(12, (9 / 4) * 3),
        strip_text_x=p9.element_text(
            margin={"t": 0.35, "b": 0.35, "l": 1, "r": 1, "units": "in"}
        ),
    )
    + p9.ylim(-1, 1)
)
```

    //anaconda3/lib/python3.7/site-packages/plotnine/layer.py:433: PlotnineWarning: position_stack : Removed 1 rows containing missing values.
      data = self.position.setup_data(self.data, params)



![png](https://drive.google.com/uc?export=view&id=16C8pj5s2DN43IKe1tKe_Onn3uSDDaIct)




After removing contributors who reported having being involved for less than a year, the overall trend holds. One concern in comparing these two datasets is that this year respondents were asked how challenging these areas were, not how much they are blocked by them. Consistent wording will make these more comparable between future surveys. 


```python
(
    make_likert_chart_multi_year(
        survey_data, "Blocker:", blocker_ratings, ["Contributing_Length", "."]
    )
    + p9.labs(
        x="Year",
        y="",
        fill="Rating",
        color="Rating",
        title="Blockers by Length of Contribution",
    )
    + p9.theme(
        figure_size=(12, (9 / 4) * 3),
        strip_text_x=p9.element_text(
            margin={"t": 0.35, "b": 0.35, "l": 1, "r": 1, "units": "in"}
        ),
        strip_text_y=p9.element_text(margin={"r": 0.8, "units": "in"}),
    )
)
```

    //anaconda3/lib/python3.7/site-packages/plotnine/layer.py:433: PlotnineWarning: position_stack : Removed 20 rows containing missing values.
      data = self.position.setup_data(self.data, params)



![png](https://drive.google.com/uc?export=view&id=1yxoeH86lX5KBbcAoJB7jwu_mD_5ngaRw)




Contributors who have been contributing for two years or more are more rate the areas as being as difficult or slightly less of an issue than in 2018. The trend is reversed for contributors with less than two years of experience. They rate items as being more difficult in 2019, suggesting potential for outreach and improvement among this group.


```python
(
    make_likert_chart_multi_year(
        survey_data, "Blocker:", blocker_ratings, ["Level_of_Contributor", "."]
    )
    + p9.labs(
        x="Blocker",
        y="Count",
        fill="Rating",
        color="Rating",
        title="Blockers by Level on Ladder",
    )
    + p9.theme(
        figure_size=(12, (9 / 4) * 3),
        strip_text_x=p9.element_text(
            margin={"t": 0.35, "b": 0.35, "l": 1, "r": 1, "units": "in"}
        ),
        strip_text_y=p9.element_text(margin={"r": 0.9, "units": "in"}),
    )
)
```

    //anaconda3/lib/python3.7/site-packages/plotnine/layer.py:433: PlotnineWarning: position_stack : Removed 7 rows containing missing values.
      data = self.position.setup_data(self.data, params)
    //anaconda3/lib/python3.7/site-packages/plotnine/layer.py:433: PlotnineWarning: position_stack : Removed 38 rows containing missing values.
      data = self.position.setup_data(self.data, params)



![png](https://drive.google.com/uc?export=view&id=1SCCqJ6PgrPfV86uS_HmHAigRnprXEJqs)




Subproject owners report less challenges in 2019 than in 2018. This is contrast to other cohorts whose challenges have increased or stayed level. This is especially clear in the columns for finding the right SIG, and code/documentation review to a lesser extent.


```python
(
    make_likert_chart_multi_year(
        survey_data, "Blocker:", blocker_ratings, ["Interested_in_next_level", "."]
    )
    + p9.labs(
        x="Blocker",
        y="Count",
        fill="Rating",
        color="Rating",
        title="Blockers by Interest in Next Level",
    )
    + p9.theme(
        figure_size=(12, (9 / 4) * 3),
        strip_text_x=p9.element_text(
            margin={"t": 0.35, "b": 0.35, "l": 1, "r": 1, "units": "in"}
        ),
        strip_text_y=p9.element_text(margin={"r": 0.9, "units": "in"}),
    )
)
```

    //anaconda3/lib/python3.7/site-packages/plotnine/layer.py:433: PlotnineWarning: position_stack : Removed 3 rows containing missing values.
      data = self.position.setup_data(self.data, params)
    //anaconda3/lib/python3.7/site-packages/plotnine/layer.py:433: PlotnineWarning: position_stack : Removed 39 rows containing missing values.
      data = self.position.setup_data(self.data, params)



![png](https://drive.google.com/uc?export=view&id=1-TMy8fI2s9LAIxdX_voZdetH9OohJMhX)



This plot shows an interesting difference between contributors interested in the next level not conditional on support, and those whose interest is conditioned on further support. In contrast to what we would expect, those who are interested irrespective of the support available report more challenges in 2019 than in 2018 with finding issues to work on and finding the right SIG. Combining this finding with the cross tabulation above showing contributors across the spectrum responding more with unconditional interest suggests several things. One is that what the mentoring program is and how it could help may need more publication. The other may be that there is a sense of pride in not needing mentoring, and so finding a way to break down that stigma may lead to happier contributors in the long run.

## Sources Checked for News


```python
make_bar_chart_multi_year(survey_data, "Check_for_news:") + p9.labs(
    title="Sources Checked for News", fill="Year", x="Source"
)
```


![png](https://drive.google.com/uc?export=view&id=1PUCic2xIHsEb15ylazuQ-D4HKQFdzD1M)



This question does not require users to choose only one option, so the proportions add up to more than 1. Most news sources see an increase in use. The largest jump is the kubernetes-dev mailing list, with over 60% of respondents checking the mailing list. 


```python
make_bar_chart_multi_year(
    survey_data, "Check_for_news:", exclude_new_contributors=True
) + p9.labs(
    title="Sources Checked for News (Excluding New Contributors)",
    fill="Year",
    x="Source",
)
```


![png](https://drive.google.com/uc?export=view&id=1db9LxC4Te4P8osmOTaIpA4nopFsIwOZF)



After excluding newer contributors, an even larger percent of contributors get their news from the mailing list, suggesting it is important for new users to be introduced to the mailing list. The other change from the previous plot is that the use of Twitter now shows a downward trend from 2018 to 2019. This confirms that advertising the survey on Twitter may have skewed the results slightly. 


```python
(
    make_bar_chart_multi_year(
        survey_data, "Check_for_news:", facet_by=[".", "Contributing_Length"]
    )
    + p9.theme(strip_text_y=p9.element_text(margin={"r": 0.8, "units": "in"}))
    + p9.labs(title="Sources Checked for News", fill="Year", x="Source", y="Count")
)
```


![png](https://drive.google.com/uc?export=view&id=1P_C7EZCJcsx57XHF0fty-p19Kal6OZG7)




In addition to the large increase in the use of Twitter by new contributors, there are a few other changes in new consumption in relation to the length of participation in Kubernetes. The proportion of new contributors using Slack has increased greatly. There is a consistent pattern across contributors who have been involved for one to two years, with many of the resources seeing decreased use. It is not clear why this is, but it not a large concern, as the most widely used resources are the same among all groups: the mailing list, Slack, and the blog. 


```python
(
    make_bar_chart_multi_year(
        survey_data, "Check_for_news:", facet_by=[".", "Interested_in_next_level"]
    )
    + p9.theme(strip_text_y=p9.element_text(margin={"r": 0.9, "units": "in"}))
    + p9.labs(title="Sources Checked for News", fill="Year", x="Source", y="Proportion")
)
```


![png](https://drive.google.com/uc?export=view&id=1PNGtcj5crAnMTOEjZ0vLrE7BSwcY0_Se)


## Resource Usage


```python
use_ratings = [
    "Every Day",
    "Several Times a Week",
    "Several Times a Month",
    "Occasionally",
    "Never",
]
use_ratings.reverse()
(
    make_likert_chart_multi_year(
        survey_data, "Use_freq:", use_ratings, five_is_high=True
    )
    + p9.labs(x="Resource", y="", fill="Rating", color="Rating", title="Resource Use")
    + p9.theme(
        figure_size=(12, (9 / 4) * 3),
        strip_text_x=p9.element_text(
            margin={"t": 0.35, "b": 0.35, "l": 1, "r": 1, "units": "in"}
        ),
        strip_text_y=p9.element_text(margin={"r": 0.9, "units": "in"}),
    )
)
```


![png](https://drive.google.com/uc?export=view&id=1WifsGFFpn--lyEZ63VjIX8iPCOMKQbQa)




The frequency of use across communication channels has fallen from 2018 to 2019 for most options. The only two with a positive trend are Slack and Google Groups/Mailing List. GitHub discussions saw a slight decrease, but is still a very heavily used tool. The decrease is due to the higher number of new contributors as shown in the plot below.


```python
use_ratings = [
    "Every Day",
    "Several Times a Week",
    "Several Times a Month",
    "Occasionally",
    "Never",
]
use_ratings.reverse()
(
    make_likert_chart_multi_year(
        survey_data,
        "Use_freq:",
        use_ratings,
        five_is_high=True,
        exclude_new_contributors=True,
    )
    + p9.labs(
        x="Resource",
        y="",
        fill="Rating",
        color="Rating",
        title="Resource Use (First Year Contributors Excluded)",
    )
    + p9.theme(
        figure_size=(12, (9 / 4) * 3),
        strip_text_x=p9.element_text(
            margin={"t": 0.35, "b": 0.35, "l": 1, "r": 1, "units": "in"}
        ),
        strip_text_y=p9.element_text(margin={"r": 0.9, "units": "in"}),
    )
)
```

![png](https://drive.google.com/uc?export=view&id=1ZsQEhse5W-Byk7k5L3TBYF6JoEgHW7Ms)




After excluding contributors with less than one year of experience, GitHub use no longer decreases. With this smaller data, we can see that out of all options and all respondents in 2018 and 2019, the use of GitHub in 2019 is the only resource to have no contributors report they never used the service.


```python
(
    make_likert_chart_multi_year(
        survey_data,
        "Use_freq:",
        use_ratings,
        facet_by=["Level_of_Contributor", "."],
        five_is_high=True,
    )
    + p9.labs(
        x="Resource",
        y="Count",
        fill="Rating",
        color="Rating",
        title="Resource Use by Contributor Level",
    )
    + p9.theme(
        figure_size=(12, (9 / 4) * 3),
        strip_text_x=p9.element_text(
            margin={"t": 0.35, "b": 0.35, "l": 1, "r": 1, "units": "in"}
        ),
        strip_text_y=p9.element_text(margin={"r": 0.9, "units": "in"}),
    )
)
```




![png](https://drive.google.com/uc?export=view&id=1m5FmKc2a32DqNzbdg_TEdz9dVJdiCE7z)




Most trends are consistent across levels of the contributor ladder. The one exception is that subproject owners, and especially reviewers, used the mailing list and Google Groups less frequently in 2019 than in 2018.

## Contribution Areas


```python
(
    make_bar_chart_multi_year(survey_data, "Contribute:")
    + p9.labs(x="Contribution", y="Count", title="Areas Contributed To", fill="Year")
)
```


![png](https://drive.google.com/uc?export=view&id=1oj9pho7UzFI1E9wuHhnKN_RUT_7U87oo)


Most areas saw increases in contributors, with the exception of code inside Kubernetes/Kubernetes and other. This reflects the continuing effort to only keep core code in the Kubernetes/Kubernetes repository, moving other contributions to additional repositories in the Kubernetes organization. 


```python
(
    make_bar_chart_multi_year(
        survey_data, "Contribute:", facet_by=["Contributing_Length", "."]
    )
    + p9.labs(
        x="Contribution", y="Proportion", title="Areas Contributed To", fill="Year"
    )
    + p9.theme(
        figure_size=(12, (9 / 4) * 3),
        strip_text_y=p9.element_text(margin={"r": 0.9, "units": "in"}),
    )
)
```


![png](https://drive.google.com/uc?export=view&id=1FQg97YCJkYs-vVhj6hqRPTxMj4L7oKsR)




Contributors who have less than one year of experience saw the greatest increase in their contributions. The largest gains were seen in the areas of Documentation and related projects. Other cohorts saw the proportions contributing to documentation slightly decrease. This isn't a bad thing if all necessary documentation is getting done, but a trend to keep an eye on if it falls below the desired standard. 


```python
(
    make_bar_chart_multi_year(
        survey_data[
            survey_data.Upstream_supported_at_employer.isin(
                ["Didn't Answer", "I’m a student"]
            )
            == False
        ],
        "Contribute:",
        facet_by=["Upstream_supported_at_employer", "."],
    )
    + p9.labs(
        x="Contribution",
        y="Count",
        title="Areas Contributed To by Employer Support",
        fill="Proportion",
    )
    + p9.theme(
        figure_size=(12, (9 / 4) * 3),
        strip_text_y=p9.element_text(margin={"r": 1, "units": "in"}),
    )
)
```

![png](https://drive.google.com/uc?export=view&id=1pRfkqZ290jrC735BjBvBKTBprpQWy3r1)


The plot above shows the proportion of each cohort that contributes to a certain area. A large drop is seen among those who use their own time in their contribution to the core repository. This is matched by an almost equal increase in the same group's contributions to other repositories owned by the Kubernetes project. 

Contributions from those who can use company time decreased in all areas. As contributors can select more than on area, this suggests that each person is contributing to less areas. This is confirmed in the table below.


```python
survey_data.groupby(["Upstream_supported_at_employer", "year"]).apply(
    lambda x: x[x.columns[x.columns.str.startswith("Contribute")]].sum(axis=1).mean()
).reset_index().rename(columns={0: "Average Number of Areas"})
```




<div>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>Upstream_supported_at_employer</th>
      <th>year</th>
      <th>Average Number of Areas</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>0</th>
      <td>Didn't Answer</td>
      <td>2018</td>
      <td>0.133333</td>
    </tr>
    <tr>
      <th>1</th>
      <td>Didn't Answer</td>
      <td>2019</td>
      <td>0.000000</td>
    </tr>
    <tr>
      <th>2</th>
      <td>I’m a student</td>
      <td>2018</td>
      <td>2.000000</td>
    </tr>
    <tr>
      <th>3</th>
      <td>it's complicated.</td>
      <td>2018</td>
      <td>1.750000</td>
    </tr>
    <tr>
      <th>4</th>
      <td>it's complicated.</td>
      <td>2019</td>
      <td>2.200000</td>
    </tr>
    <tr>
      <th>5</th>
      <td>no, I need to use my own time</td>
      <td>2018</td>
      <td>1.833333</td>
    </tr>
    <tr>
      <th>6</th>
      <td>no, I need to use my own time</td>
      <td>2019</td>
      <td>1.896552</td>
    </tr>
    <tr>
      <th>7</th>
      <td>yes, I can contribute on company time</td>
      <td>2018</td>
      <td>2.812500</td>
    </tr>
    <tr>
      <th>8</th>
      <td>yes, I can contribute on company time</td>
      <td>2019</td>
      <td>2.551724</td>
    </tr>
  </tbody>
</table>
</div>




```python
(
    make_bar_chart_multi_year(
        survey_data, "Contribute:", facet_by=["Interested_in_next_level", "."]
    )
    + p9.labs(
        x="Contribution",
        y="Count",
        title="Areas Contributed To by Interest in Next Level",
        fill="Year",
    )
    + p9.theme(
        figure_size=(12, (9 / 4) * 3),
        strip_text_y=p9.element_text(margin={"r": 1, "units": "in"}),
    )
)
```


![png](https://drive.google.com/uc?export=view&id=1comh4Ip4ehEjz-qXTVuBctZo8-Egsnp4)



Core code contributions dropped from 2018 to 2019 among all groups except subproject owners. Those who want support to reach the next level contributed more across all areas. Those who want to move to the next level but do not condition their interest on receiving help contributed less across most areas than in 2018. This again suggests that more clarification is needed around the mentorship program, as many of those who think they need more support contribute widely, while those that don't contribute as much don't feel they need support. 

## Use of Help Wanted and Good First Issue Labels


```python
help_wanted = survey_data[
    survey_data[
        "Do_you_use_the\xa0Help_Wanted_and/or_Good_First_Issue_labels_on_issues_you_file_to_find_contributors"
    ].isna()
    == False
]
```


```python
help_counts = (
    help_wanted.rename(
        columns={
            "Do_you_use_the\xa0Help_Wanted_and/or_Good_First_Issue_labels_on_issues_you_file_to_find_contributors": "help"
        }
    )
    .groupby(["year", "help"])
    .count()["Collector_ID"]
    .reset_index()
)
```


```python
help_year_counts = help_counts.groupby("year").sum().reset_index()
help_percents = help_year_counts.merge(help_counts, on="year")
```


```python
help_percents = help_percents.assign(
    percent=help_percents["Collector_ID_y"] / help_percents["Collector_ID_x"]
)
```


```python
help_plot = (
    p9.ggplot(help_percents, p9.aes(x="help", y="percent", fill="factor(year)"))
    + p9.geom_bar(position=p9.position_dodge(preserve="single"), stat="identity")
    + p9.theme(axis_text_x=p9.element_text(angle=45, ha="right"))
    + p9.labs(
        x="Used Label",
        title="Use of Help Wanted and/or Good First Issue Labels",
        y="Proportion",
        fill="Year",
    )
)

help_plot
```


![png](https://drive.google.com/uc?export=view&id=1lq4vReECvnMV4-llylvgUVv6JS_nYxvh)


Use of the help wanted and good first labels clearly increased from 2018 to 2019.
