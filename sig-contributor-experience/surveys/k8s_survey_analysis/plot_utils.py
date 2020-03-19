from textwrap import wrap
import math
import plotnine as p9
import pandas as pd
import textwrap
from textwrap import shorten
from matplotlib import pyplot as plt
from copy import copy


from mizani.palettes import brewer_pal
from plotnine.scales.scale import scale_discrete

# Custom scales for plotnine that reverse the direction of the colors
class reverse_scale_color_brewer(p9.scale_color_brewer):
    def __init__(self, type="seq", palette=1, direction=-1, **kwargs):
        self.palette = brewer_pal(type, palette, direction)
        scale_discrete.__init__(self, **kwargs)


class reverse_scale_fill_brewer(p9.scale_fill_brewer):
    def __init__(self, type="seq", palette=1, direction=-1, **kwargs):
        self.palette = brewer_pal(type, palette, direction)
        scale_discrete.__init__(self, **kwargs)


def split_for_likert(topic_data_long, mid_point):
    """
    Returns the aggregated counts for ratings in the top and bottom halves of 
    the of each category, necssary for making offset bar charts

    Args:
        topic_data_long (pandas.Dataframe): A pandas Dataframe storing each respondents 
        ratings for a given topic, in long format
        mid_point (int): The midpoint to use to split the into two halves, based on ratings

    Returns:
        (tuple): Tuple containing:
            (pandas.DataFrame): Aggregated counts for ratings greater than or equal to the midpoinnt
            (pandas.DataFrame): Aggregated counts for ratings less than or equal to the midpoinnt 
    """
    x = topic_data_long.columns.tolist()
    x.remove("level_1")

    top_cutoff = topic_data_long["rating"] >= mid_point
    bottom_cutoff = topic_data_long["rating"] <= mid_point

    top_scores = (
        topic_data_long[top_cutoff]
        .groupby(x)
        .count()
        .reindex(
            pd.MultiIndex.from_product(
                [topic_data_long[y].unique().tolist() for y in x], names=x
            ),
            fill_value=0,
        )
        .reset_index()
        .sort_index(ascending=False)
    )

    # The mid point is in both the top and bottom halves, so divide by two
    top_scores.loc[top_scores["rating"] == mid_point, "level_1"] = (
        top_scores[top_scores["rating"] == mid_point]["level_1"] / 2.0
    )

    bottom_scores = (
        topic_data_long[bottom_cutoff]
        .groupby(x)
        .count()
        .reindex(
            pd.MultiIndex.from_product(
                [topic_data_long[y].unique().tolist() for y in x], names=x
            ),
            fill_value=0,
        )
        .reset_index()
    )

    # The mid point is in both the top and bottom halves, so divide by two
    bottom_scores.loc[bottom_scores["rating"] == mid_point, "level_1"] = (
        bottom_scores[bottom_scores["rating"] == mid_point]["level_1"] / 2.0
    )

    return top_scores, bottom_scores


def make_long(data, facets, multi_year=False):
    """Converts a wide dataframe with columns for each topic's rating into a long dataframe

    Args:
        data (pandas.DataFrame): A wide dataframe
        facets (list): List of columns to keep as their own column
        mulit_year (bool, optional) Defaults to False. If True, add the "year" column to the list of facets

    Returns:
        (pandas.DataFrame): Long dataframe 

    """

    facets = copy(facets)
    if multi_year:
        facets.append("year")
    long_data = data.set_index(facets, append=True).stack().reset_index()

    # Rename so Level_0 always has the values of the topic we are interested in
    long_data = long_data.rename(
        columns={
            "level_0": "level_1",
            "level_4": "level_0",
            "level_3": "level_0",
            "level_2": "level_0",
            0: "rating",
        }
    )
    long_data = long_data.assign(
        level_0=pd.Categorical(long_data.level_0, ordered=True)
    )
    return long_data


def get_data_subset(
    survey_data, topic, facets=[], exclude_new_contributors=False, include_year=False
):
    """Get only the relevant columns from the data

    Args:
        survey_data (pandas.DataFrame): Raw data read in from Kubernetes Survey   
        topic (str): String that all questions of interest start with 
        facets (list, optional): List of columns use for grouping
        exclude_new_contributors: (bool, optional) Defaults to False. If True, remove 
            all responses from contributors who have been involved a year or less.  
        include_year: (bool, optional) Defaults to False. If True, include the year column
            in the output

    Returns:
        (pandas.DataFrame): Survey dataframe with only columns relevant to the topics
            and facets remaining.
    """

    og_cols = [x for x in survey_data.columns if x.startswith(topic)]
    facets = copy(facets)
    if include_year:
        facets.append("year")
    if facets:
        if "." in facets:
            facets.remove(".")
            cols = og_cols + facets
            facets.append(".")
        else:
            cols = og_cols + facets
    else:
        cols = og_cols

    if exclude_new_contributors:
        topic_data = survey_data[
            survey_data["Contributing_Length"] != "less than one year"
        ][cols]
    else:
        topic_data = survey_data[cols]

    return topic_data


def get_multi_year_data_subset(
    survey_data, topic, facet_by=[], exclude_new_contributors=False
):
    """Get appropriate data for multi-year plots and convert it to long form

    Args:
        survey_data (pandas.DataFrame): Raw data read in from Kubernetes Survey   
        topic (str): String that all questions of interest start with 
        facet_by (list, optional): List of columns use for grouping
        exclude_new_contributors (bool, optional) Defaults to False. If True, remove 
            all responses from contributors who have been involved a year or less.  

    Returns:
        (pandas.DataFrame): Long dataframe 
    """
    topic_data = get_data_subset(
        survey_data, topic, facet_by, exclude_new_contributors, include_year=True
    )

    if facet_by:
        if "." in facet_by:
            facet_by.remove(".")
            topic_data_long = make_long(topic_data, facet_by, multi_year=True)
            facet_by.append(".")
        else:
            topic_data_long = make_long(topic_data, facet_by, multi_year=True)

    else:
        topic_data_long = make_long(topic_data, [], multi_year=True)

    return topic_data_long


def get_single_year_data_subset(survey_data, topic, facet_by=[]):
    """Get appropriate data for single-year plots and convert it to long form

    Args:
        survey_data (pandas.DataFrame): Raw data read in from Kubernetes Survey   
        topic (str): String that all questions of interest start with 
        facet_by (list, optional): List of columns use for grouping

    Returns:
        (pandas.DataFrame): Long dataframe 

    """
    topic_data = get_data_subset(survey_data, topic, facet_by)

    if facet_by:
        if "." in facet_by:
            facet_by.remove(".")
            topic_data_long = make_long(topic_data, facet_by)
            facet_by.append(".")
        else:
            topic_data_long = make_long(topic_data, facet_by)
    else:

        topic_data_long = (
            topic_data.unstack().reset_index().rename(columns={0: "rating"})
        )
        topic_data_long = topic_data_long.assign(
            level_0=pd.Categorical(topic_data_long.level_0, ordered=True)
        )

    return topic_data_long


def make_bar_chart_multi_year(
    survey_data, topic, facet_by=[], exclude_new_contributors=False
):
    """Make a barchart showing proportions of respondents listing each 
        column that starts with topic. Bars are colored by which year of 
        the survey they correspond to. If facet_by is not empty, the resulting
        plot will be faceted into subplots by the variables given. 

    Args:
        survey_data (pandas.DataFrame): Raw data read in from Kubernetes Survey   
        topic (str): String that all questions of interest start with 
        facet_by (list,optional): List of columns use for grouping
        exclude_new_contributors (bool, optiona ): Defaults to False. If True,
            do not include any responses from contributors with less than 
            one year of experience

    Returns:
        (plotnine.ggplot): Plot object which can be displayed in a notebook or saved out to a file

    """
    topic_data = get_data_subset(
        survey_data, topic, facet_by, exclude_new_contributors, include_year=True
    )

    if facet_by:
        fix = False
        if "." in facet_by:
            facet_by.remove(".")
            fix = True
        agg = (
            topic_data.groupby(facet_by + ["year"])
            .sum()
            .reset_index()
            .melt(id_vars=facet_by + ["year"])
        )
        totals = (
            topic_data.groupby(facet_by + ["year"])
            .count()
            .reset_index()
            .melt(id_vars=facet_by + ["year"])
        )
        percent = agg.merge(totals, on=facet_by + ["year", "variable"])

        if fix:
            facet_by.append(".")

    else:
        agg = topic_data.groupby(["year"]).sum().reset_index().melt(id_vars=["year"])
        totals = (
            topic_data.groupby(["year"]).count().reset_index().melt(id_vars=["year"])
        )
        percent = agg.merge(totals, on=["year", "variable"])

    # This plot is always done proportionally
    percent = percent.assign(value=percent["value_x"] / percent["value_y"])
    percent = percent.assign(variable=pd.Categorical(percent.variable, ordered=True))

    br = (
        p9.ggplot(percent, p9.aes(x="variable", fill="factor(year)", y="value"))
        + p9.geom_bar(show_legend=True, position="dodge", stat="identity")
        + p9.theme(
            axis_text_x=p9.element_text(angle=45, ha="right"),
            strip_text_y=p9.element_text(angle=0, ha="left"),
        )
        + p9.scale_x_discrete(
            limits=sorted(percent["variable"].unique().tolist()),
            labels=[
                shorten(
                    x.replace(topic, "").replace("_", " "), placeholder="...", width=30
                )
                for x in sorted(percent["variable"].unique().tolist())
            ],
        )
    )

    # Uncomment to return dataframe instead of plot
    # return percent

    if facet_by:
        br = (
            br
            + p9.facet_grid(
                facet_by,
                shrink=False,
                labeller=lambda x: "\n".join(wrap(x.replace("/", "/ "), 15)),
            )
            + p9.theme(
                strip_text_x=p9.element_text(wrap=True, va="bottom", margin={"b": -0.5})
            )
        )
    return br


def make_single_bar_chart_multi_year(survey_data, column, facet, proportionally=False):
    """Make a barchart showing the number of respondents responding to a single column.
        Bars are colored by which year of the survey they correspond to. If facet
        is not empty, the resulting plot will be faceted into subplots by the variables
        given. 

    Args:
        survey_data (pandas.DataFrame): Raw data read in from Kubernetes Survey   
        column (str): Column to plot responses to
        facet (list,optional): List of columns use for grouping
        proportionally (bool, optiona ): Defaults to False. If True,
            the bars heights are determined proportionally to the 
            total number of responses in that facet. 

    Returns:
        (plotnine.ggplot): Plot object which can be displayed in a notebook or saved out to a file

    """
    cols = [column, facet]
    show_legend = False
    topic_data = survey_data[cols + ["year"]]

    topic_data_long = make_long(topic_data, facet, multi_year=True)

    if proportionally:
        proportions = (
            topic_data_long[topic_data_long.rating == 1].groupby(facet + ["year"]).sum()
            / topic_data_long.groupby(facet + ["year"]).sum()
        ).reset_index()
    else:
        proportions = (
            topic_data_long[topic_data_long.rating == 1]
            .groupby(facet + ["year"])
            .count()
            .reset_index()
        )

    x = topic_data_long.columns.tolist()
    x.remove("level_1")

    ## Uncomment to return dataframe instead of plot
    # return proportions

    return (
        p9.ggplot(proportions, p9.aes(x=facet, fill="year", y="level_1"))
        + p9.geom_bar(show_legend=show_legend, stat="identity")
        + p9.theme(
            axis_text_x=p9.element_text(angle=45, ha="right"),
            strip_text_y=p9.element_text(angle=0, ha="left"),
        )
        + p9.scale_x_discrete(
            limits=topic_data_long[facet].unique().tolist(),
            labels=[
                x.replace("_", " ") for x in topic_data_long[facet].unique().tolist()
            ],
        )
    )


def make_likert_chart_multi_year(
    survey_data,
    topic,
    labels,
    facet_by=[],
    five_is_high=False,
    exclude_new_contributors=False,
):
    """Make an offset stacked barchart showing the number of respondents at each rank or value for 
        all columns in the topic. Each column in the topic is a facet, with the years displayed
        along the x-axis.

    Args:
        survey_data (pandas.DataFrame): Raw data read in from Kubernetes Survey   
        topic (str): String that all questions of interest start with
        labels (list): List of strings to use as labels, corresponding
             to the numerical values given by the respondents.
        facet_by (list,optional): List of columns use for grouping
        five_is_high (bool, optiona ): Defaults to False. If True,
            five is considered the highest value in a ranking, otherwise 
            it is taken as the lowest value.
        exclude_new_contributors (bool, optional): Defaults to False. If True,
            do not include any responses from contributors with less than 
            one year of experience        

    Returns:
        (plotnine.ggplot): Offset stacked barchart plot object which 
            can be displayed in a notebook or saved out to a file
    """

    facet_by = copy(facet_by)
    og_cols = [x for x in survey_data.columns if x.startswith(topic)]
    show_legend = True

    topic_data_long = get_multi_year_data_subset(
        survey_data, topic, facet_by, exclude_new_contributors
    )

    if not five_is_high:
        topic_data_long = topic_data_long.assign(rating=topic_data_long.rating * -1.0)

    mid_point = 3 if five_is_high else -3
    top_scores, bottom_scores = split_for_likert(topic_data_long, mid_point)

    if facet_by:
        fix = False
        if "." in facet_by:
            facet_by.remove(".")
            fix = True

        # Calculate proportion for each rank
        top_scores = top_scores.merge(
            topic_data_long.groupby(facet_by + ["year"]).count().reset_index(),
            on=facet_by + ["year"],
        ).rename(columns={"rating_x": "rating", "level_0_x": "level_0"})
        top_scores = top_scores.assign(
            level_1=top_scores.level_1_x / (top_scores.level_1_y / len(og_cols))
        )

        bottom_scores = bottom_scores.merge(
            topic_data_long.groupby(facet_by + ["year"]).count().reset_index(),
            on=facet_by + ["year"],
        ).rename(columns={"rating_x": "rating", "level_0_x": "level_0"})
        bottom_scores = bottom_scores.assign(
            level_1=bottom_scores.level_1_x
            * -1
            / (bottom_scores.level_1_y / len(og_cols))
        )

        if fix:
            facet_by.append(".")
    else:
        # Calculate proportion for each rank
        top_scores = top_scores.merge(
            topic_data_long.groupby(["year"]).count().reset_index(), on=["year"]
        ).rename(columns={"rating_x": "rating", "level_0_x": "level_0"})
        top_scores = top_scores.assign(
            level_1=top_scores.level_1_x / (top_scores.level_1_y / len(og_cols))
        )

        bottom_scores = bottom_scores.merge(
            topic_data_long.groupby(["year"]).count().reset_index(), on=["year"]
        ).rename(columns={"rating_x": "rating", "level_0_x": "level_0"})
        bottom_scores = bottom_scores.assign(
            level_1=bottom_scores.level_1_x
            * -1
            / (bottom_scores.level_1_y / len(og_cols))
        )

    vp = (
        p9.ggplot(
            topic_data_long,
            p9.aes(x="factor(year)", fill="factor(rating)", color="factor(rating)"),
        )
        + p9.geom_col(
            data=top_scores,
            mapping=p9.aes(y="level_1"),
            show_legend=show_legend,
            size=0.25,
            position=p9.position_stack(reverse=True),
        )
        + p9.geom_col(
            data=bottom_scores,
            mapping=p9.aes(y="level_1"),
            show_legend=show_legend,
            size=0.25,
            position=p9.position_stack(),
        )
        + p9.geom_hline(yintercept=0, color="white")
    )

    if five_is_high:
        vp = (
            vp
            + p9.scale_color_brewer(
                "div", "RdBu", limits=[1, 2, 3, 4, 5], labels=labels
            )
            + p9.scale_fill_brewer("div", "RdBu", limits=[1, 2, 3, 4, 5], labels=labels)
            + p9.theme(
                axis_text_x=p9.element_text(angle=45, ha="right"),
                strip_text_y=p9.element_text(angle=0, ha="left"),
            )
        )
    else:
        vp = (
            vp
            + p9.scale_color_brewer(
                "div", "RdBu", limits=[-5, -4, -3, -2, -1], labels=labels
            )
            + p9.scale_fill_brewer(
                "div", "RdBu", limits=[-5, -4, -3, -2, -1], labels=labels
            )
            + p9.theme(strip_text_y=p9.element_text(angle=0, ha="left"))
        )

    if facet_by:
        facet_by.remove(".")

    else:
        facet_by.append(".")

    vp = (
        vp
        + p9.facet_grid(
            facet_by + ["level_0"],
            labeller=lambda x: "\n".join(
                wrap(
                    x.replace(topic, "").replace("_", " ").replace("/", "/ ").strip(),
                    15,
                )
            ),
        )
        + p9.theme(
            strip_text_x=p9.element_text(wrap=True, ma="left"), panel_spacing_x=0.1
        )
    )

    return vp


def make_bar_chart(survey_data, topic, facet_by=[], proportional=False):
    """Make a barchart showing the number of respondents listing each 
        column that starts with topic for a single year. If facet_by is
        not empty, the resulting plot will be faceted into subplots 
        by the variables given. 

    Args:
        survey_data (pandas.DataFrame): Raw data read in from Kubernetes Survey   
        topic (str): String that all questions of interest start with 
        facet_by (list,optional): List of columns use for grouping
        proportional (bool, optiona ): Defaults to False. If True,
            the bars heights are determined proportionally to the 
            total number of responses in that facet. 

    Returns:
        (plotnine.ggplot): Plot object which can be displayed in a notebook or saved out to a file
    """
    show_legend = False
    if facet_by:
        show_legend = True

    topic_data_long = get_single_year_data_subset(survey_data, topic, facet_by)

    x = topic_data_long.columns.tolist()
    x.remove("level_1")

    if facet_by:
        period = False
        if "." in facet_by:
            facet_by.remove(".")
            period = True

        aggregate_data = (
            topic_data_long[topic_data_long.rating == 1]
            .dropna()
            .groupby(["level_0"] + facet_by)
            .count()
            .reset_index()
        )

        if period:
            facet_by.append(".")

    else:
        aggregate_data = (
            topic_data_long[topic_data_long.rating == 1]
            .dropna()
            .groupby("level_0")
            .count()
            .reset_index()
        )

    if proportional and facet_by:
        period = False
        if "." in facet_by:
            facet_by.remove(".")
            period = True

        facet_sums = (
            topic_data_long[topic_data_long.rating == 1]
            .dropna()
            .groupby(facet_by)
            .count()
            .reset_index()
        )

        aggregate_data = aggregate_data.merge(facet_sums, on=facet_by).rename(
            columns={"level_0_x": "level_0"}
        )
        aggregate_data = aggregate_data.assign(
            rating=aggregate_data.rating_x / aggregate_data.rating_y
        )

        if period:
            facet_by.append(".")

    br = (
        p9.ggplot(aggregate_data, p9.aes(x="level_0", fill="level_0", y="rating"))
        + p9.geom_bar(show_legend=show_legend, stat="identity")
        + p9.theme(
            axis_text_x=p9.element_text(angle=45, ha="right"),
            strip_text_y=p9.element_text(angle=0, ha="left"),
        )
        + p9.scale_x_discrete(
            limits=topic_data_long["level_0"].unique().tolist(),
            labels=[
                "\n".join(
                    textwrap.wrap(x.replace(topic, "").replace("_", " "), width=35)[0:2]
                )
                for x in topic_data_long["level_0"].unique().tolist()
            ],
        )
    )

    if facet_by:
        br = (
            br
            + p9.facet_grid(
                facet_by, shrink=False, labeller=lambda x: "\n".join(wrap(x, 15))
            )
            + p9.theme(
                axis_text_x=p9.element_blank(),
                strip_text_x=p9.element_text(
                    wrap=True, va="bottom", margin={"b": -0.5}
                ),
            )
            + p9.scale_fill_discrete(
                limits=topic_data_long["level_0"].unique().tolist(),
                labels=[
                    "\n".join(
                        wrap(
                            x.replace(topic, "")
                            .replace("_", " ")
                            .replace("/", "/  ")
                            .strip(),
                            30,
                        )
                    )
                    for x in topic_data_long["level_0"].unique().tolist()
                ],
            )
        )
    return br


def make_likert_chart(
    survey_data,
    topic,
    labels,
    facet_by=[],
    max_value=5,
    max_is_high=False,
    wrap_facets=True,
    sort_x=False,
):
    """Make an offset stacked barchart showing the number of respondents at each rank or value for 
        all columns in the topic. Each column in the original data is a tick on the x-axis

    Args:
        survey_data (pandas.DataFrame): Raw data read in from Kubernetes Survey   
        topic (str): String that all questions of interest start with
        labels (list): List of strings to use as labels, corresponding
             to the numerical values given by the respondents.
        facet_by (list,optional): List of columns use for grouping 
        max_value (int, optional):  Defaults to 5. The maximuum value a respondent can assign.
        max_is_high (bool, optiona ): Defaults to False. If True,
            the max_value is considered the highest value in a ranking, otherwise 
            it is taken as the lowest value.
        wrap_facets (bool, optional): Defaults to True. If True, the facet labels are 
            wrapped
        sort_x  (bool, optional): Defaults to False. If True, the x-axis is sorted by the 
            mean value for each column in the original data 

    Returns:
        (plotnine.ggplot): Offset stacked barchart plot object which 
            can be displayed in a notebook or saved out to a file
    """

    mid_point = math.ceil(max_value / 2)

    og_cols = [x for x in survey_data.columns if x.startswith(topic)]
    show_legend = True

    topic_data_long = get_single_year_data_subset(survey_data, topic, facet_by)

    if not max_is_high:
        topic_data_long = topic_data_long.assign(rating=topic_data_long.rating * -1.0)

        mid_point = -1 * mid_point

    top_scores, bottom_scores = split_for_likert(topic_data_long, mid_point)

    if facet_by:
        fix = False
        if "." in facet_by:
            facet_by.remove(".")
            fix = True

        top_scores = top_scores.merge(
            topic_data_long.groupby(facet_by).count().reset_index(), on=facet_by
        ).rename(columns={"rating_x": "rating", "level_0_x": "level_0"})
        top_scores = top_scores.assign(
            level_1=top_scores.level_1_x / (top_scores.level_1_y / len(og_cols))
        )

        bottom_scores = bottom_scores.merge(
            topic_data_long.groupby(facet_by).count().reset_index(), on=facet_by
        ).rename(columns={"rating_x": "rating", "level_0_x": "level_0"})
        bottom_scores = bottom_scores.assign(
            level_1=bottom_scores.level_1_x
            * -1
            / (bottom_scores.level_1_y / len(og_cols))
        )

        if fix:
            facet_by.append(".")

    else:
        bottom_scores = bottom_scores.assign(level_1=bottom_scores.level_1 * -1)

    if sort_x:
        x_sort_order = (
            topic_data_long.groupby("level_0")
            .mean()
            .sort_values("rating")
            .reset_index()["level_0"]
            .values.tolist()
        )
        x_sort_order.reverse()
    else:
        x_sort_order = topic_data_long["level_0"].unique().tolist()

    vp = (
        p9.ggplot(
            topic_data_long,
            p9.aes(x="level_0", fill="factor(rating)", color="factor(rating)"),
        )
        + p9.geom_col(
            data=top_scores,
            mapping=p9.aes(y="level_1"),
            show_legend=show_legend,
            size=0.25,
            position=p9.position_stack(reverse=True),
        )
        + p9.geom_col(
            data=bottom_scores,
            mapping=p9.aes(y="level_1"),
            show_legend=show_legend,
            size=0.25,
            position=p9.position_stack(),
        )
        + p9.geom_hline(yintercept=0, color="white")
        + p9.theme(
            axis_text_x=p9.element_text(angle=45, ha="right"),
            strip_text_y=p9.element_text(angle=0, ha="left"),
        )
        + p9.scale_x_discrete(
            limits=x_sort_order,
            labels=[
                "\n".join(
                    textwrap.wrap(x.replace(topic, "").replace("_", " "), width=35)[0:2]
                )
                for x in x_sort_order
            ],
        )
    )

    if max_is_high:
        vp = (
            vp
            + p9.scale_color_brewer(
                "div", "RdBu", limits=list(range(1, max_value + 1)), labels=labels
            )
            + p9.scale_fill_brewer(
                "div", "RdBu", limits=list(range(1, max_value + 1)), labels=labels
            )
        )

    else:
        vp = (
            vp
            + reverse_scale_fill_brewer(
                "div",
                "RdBu",
                limits=list(reversed(range(-max_value, 0))),
                labels=labels,
            )
            + reverse_scale_color_brewer(
                "div",
                "RdBu",
                limits=list(reversed(range(-max_value, 0))),
                labels=labels,
            )
        )

    if facet_by:
        if wrap_facets:
            vp = (
                vp
                + p9.facet_grid(facet_by, labeller=lambda x: "\n".join(wrap(x, 15)))
                + p9.theme(
                    strip_text_x=p9.element_text(
                        wrap=True, va="bottom", margin={"b": -0.5}
                    )
                )
            )
        else:
            vp = vp + p9.facet_grid(facet_by, space="free", labeller=lambda x: x)
    return vp


def make_single_likert_chart(survey_data, column, facet, labels, five_is_high=False):
    """Make an offset stacked barchart showing the number of respondents at each rank 
        or value for a single columns in the original data. Each facet is shown as
        a tick on the x-axis

    Args:
        survey_data (pandas.DataFrame): Raw data read in from Kubernetes Survey   
        topic (str): String that all questions of interest start with
        labels (list): List of strings to use as labels, corresponding
             to the numerical values given by the respondents.
        facet (str): Column used for grouping 
        five_is_high (bool, optionalc): Defaults to False. If True,
            5 is considered the highest value in a ranking, otherwise 
            it is taken as the lowest value.

    Returns:
        (plotnine.ggplot): Offset stacked barchart plot object which 
            can be displayed in a notebook or saved out to a file
    """
    mid_point = 3
    cols = [column, facet]
    show_legend = True
    topic_data = survey_data[cols]

    topic_data_long = make_long(topic_data, facet)

    if not five_is_high:
        topic_data_long = topic_data_long.assign(rating=topic_data_long.rating * -1.0)
    x = topic_data_long.columns.tolist()
    x.remove("level_1")
    x.remove("level_0")

    if not five_is_high:
        mid_point *= -1

    top_cutoff = topic_data_long["rating"] >= mid_point
    bottom_cutoff = topic_data_long["rating"] <= mid_point

    top_scores = (
        topic_data_long[top_cutoff]
        .groupby(x)
        .count()
        .reset_index()
        .sort_index(ascending=False)
    )

    top_scores.loc[top_scores["rating"] == mid_point, "level_1"] = (
        top_scores[top_scores["rating"] == mid_point]["level_1"] / 2.0
    )
    top_scores = top_scores.merge(
        topic_data_long.groupby(facet).count().reset_index(), on=facet
    )
    top_scores = top_scores.assign(level_1=top_scores.level_1_x / top_scores.level_1_y)

    bottom_scores = topic_data_long[bottom_cutoff].groupby(x).count().reset_index()
    bottom_scores.loc[bottom_scores["rating"] == mid_point, "level_1"] = (
        bottom_scores[bottom_scores["rating"] == mid_point]["level_1"] / 2.0
    )
    bottom_scores = bottom_scores.merge(
        topic_data_long.groupby(facet).count().reset_index(), on=facet
    )
    bottom_scores = bottom_scores.assign(
        level_1=bottom_scores.level_1_x * -1 / bottom_scores.level_1_y
    )

    vp = (
        p9.ggplot(
            topic_data_long,
            p9.aes(x=facet, fill="factor(rating_x)", color="factor(rating_x)"),
        )
        + p9.geom_col(
            data=top_scores,
            mapping=p9.aes(y="level_1"),
            show_legend=show_legend,
            size=0.25,
            position=p9.position_stack(reverse=True),
        )
        + p9.geom_col(
            data=bottom_scores,
            mapping=p9.aes(y="level_1"),
            show_legend=show_legend,
            size=0.25,
        )
        + p9.geom_hline(yintercept=0, color="white")
        + p9.theme(
            axis_text_x=p9.element_text(angle=45, ha="right"),
            strip_text_y=p9.element_text(angle=0, ha="left"),
        )
        + p9.scale_x_discrete(
            limits=topic_data_long[facet].unique().tolist(),
            labels=[
                x.replace("_", " ") for x in topic_data_long[facet].unique().tolist()
            ],
        )
    )

    if five_is_high:
        vp = (
            vp
            + p9.scale_color_brewer(
                "div",
                "RdBu",
                limits=[1, 2, 3, 4, 5],
                labels=["\n".join(wrap(x, 15)) for x in labels],
            )
            + p9.scale_fill_brewer(
                "div",
                "RdBu",
                limits=[1, 2, 3, 4, 5],
                labels=["\n".join(wrap(x, 15)) for x in labels],
            )
        )
    else:
        vp = (
            vp
            + reverse_scale_fill_brewer(
                "div",
                "RdBu",
                limits=[-1, -2, -3, -4, -5],
                labels=["\n".join(wrap(x, 15)) for x in labels],
            )
            + reverse_scale_color_brewer(
                "div",
                "RdBu",
                limits=[-1, -2, -3, -4, -5],
                labels=["\n".join(wrap(x, 15)) for x in labels],
            )
        )

    return vp


def make_single_bar_chart(
    survey_data, column, facet, proportionally=False, facet2=None
):
    """Make a barchart showing the number of respondents marking 
        a certain column in the original dataset as True. The facet
        variable values are used as ticks on the x-axis

    Args:
        survey_data (pandas.DataFrame): Raw data read in from Kubernetes Survey   
        topic (str): String that all questions of interest start with 
        facet (str): Column use for grouping
        proportional (bool, optiona ): Defaults to False. If True,
            the bars heights are determined proportionally to the 
            total number of responses in that facet. 
        facet2 (str, optional): If provided, a second variable to facet against.

    Returns:
        (plotnine.ggplot): Plot object which can be displayed in a notebook or saved out to a file
    """
    cols = [column, facet]
    if facet2:
        cols.append(facet2)
    show_legend = False
    topic_data = survey_data[cols]

    grouper = [facet, facet2] if facet2 else facet
    topic_data_long = make_long(topic_data, grouper)

    if proportionally:
        proportions = (
            topic_data_long[topic_data_long.rating == 1].groupby(grouper).sum()
            / topic_data_long.groupby(grouper).sum()
        ).reset_index()
    else:
        proportions = (
            topic_data_long[topic_data_long.rating == 1]
            .groupby(grouper)
            .count()
            .reset_index()
        )

    x = topic_data_long.columns.tolist()
    x.remove("level_1")

    br = (
        p9.ggplot(proportions, p9.aes(x=facet, fill=facet, y="level_1"))
        + p9.geom_bar(show_legend=show_legend, stat="identity")
        + p9.theme(
            axis_text_x=p9.element_text(angle=45, ha="right"),
            strip_text_y=p9.element_text(angle=0, ha="left"),
        )
        + p9.scale_x_discrete(
            limits=topic_data_long[facet].unique().tolist(),
            labels=[
                x.replace("_", " ") for x in topic_data_long[facet].unique().tolist()
            ],
        )
    )

    if facet2:
        br = br + p9.facet_grid([facet2, "."])

    return br
