# Resources

## Kubernetes integration status by big data product

### Spark

[Apache Spark](https://spark.apache.org) is a distributed data processing framework. 

##### Status

Kubernetes is supported as a mainline Spark scheduler since [release 2.3](https://spark.apache.org/releases/spark-release-2-3-0.html), see [the detailed documentation](https://spark.apache.org/docs/latest/running-on-kubernetes.html).
That work was done after the [Spark on Kubernetes original Design Proposal](https://docs.google.com/document/d/1_bBzOZ8rKiOSjQg78DXOA3ZBIo_KkDJjqxVuq0yXdew/edit#)
in the [apache-spark-on-k8s git repo](https://github.com/apache-spark-on-k8s/spark).

##### Activities 

Enhancements are under development, with a good overview given [in this blog post](https://databricks.com/blog/2018/09/26/whats-new-for-apache-spark-on-kubernetes-in-the-upcoming-apache-spark-2-4-release.html).

* Work is underway for Spark 2.4 to improve support and integration with HDFS.
  * Design Document: [How Spark on Kubernetes will access Secure HDFS](https://docs.google.com/document/d/1RBnXD9jMDjGonOdKJ2bA1lN4AAV_1RwpU_ewFuCNWKg/edit#heading=h.verdza2f4fyd)
* Shuffle service design
  * Design Document [Improving Spark Shuffle Reliability](https://docs.google.com/document/d/1uCkzGGVG17oGC6BJ75TpzLAZNorvrAU3FRd2X-rVHSM/edit)
  * JIRA issue [SPARK-25299: Use remote storage for persisting shuffle data](https://issues.apache.org/jira/browse/SPARK-25299)

### HDFS

[Apache Hadoop HDFS](https://hadoop.apache.org/hdfs) is a distributed file system, the persistence layer for Hadoop.

##### Status

TODO, e.g. "No release yet."

##### Activities

* [Data Locality Doc](https://docs.google.com/document/d/1TAC6UQDS3M2sin2msFcZ9UBBQFyyz4jFKWw5BM54cQo/edit)
* ["HDFS on Kubernetes" git repository including Helm charts](https://github.com/apache-spark-on-k8s/kubernetes-HDFS)

### Airflow

[Apache Airflow](https://airflow.apache.org) is a platform to programmatically author, schedule and monitor workflows.

##### Status

The [Kubernetes executor](https://airflow.apache.org/kubernetes.html)  has been introduced with Airflow [release 1.10.0](https://github.com/apache/incubator-airflow/blob/master/CHANGELOG.txt)  with support of Kubernetes 1.10. 

##### Activities

* [Airflow roadmap](https://cwiki.apache.org/confluence/pages/viewpage.action?pageId=71013666)

### Flink

[Apache Flink](https://flink.apache.org) is a distributed data processing framework.

##### Status

Flink 1.6 supports [running a session or job cluster on Kubernetes](https://ci.apache.org/projects/flink/flink-docs-stable/ops/deployment/kubernetes.html).

##### Activities

* [Native support for Kubernetes as a Flink runtime](https://issues.apache.org/jira/browse/FLINK-9953) 
* [Lyft is working on an operator](https://lists.apache.org/thread.html/aa941030440c1d9e34c35c0caf5ddd2456755337fc34a4edebb32929@%3Cdev.flink.apache.org%3E)
