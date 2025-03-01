package google

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"testing"
	"time"

	// "regexp"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	dataproc "google.golang.org/api/dataproc/v1beta2"
	"google.golang.org/api/googleapi"
)

type jobTestField struct {
	tf_attr  string
	gcp_attr interface{}
}

// TODO (mbang): Test `ExactlyOneOf` here
// func TestAccDataprocJob_failForMissingJobConfig(t *testing.T) {
// 	t.Parallel()

// 	vcrTest(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccCheckDataprocJobDestroyProducer(t),
// 		Steps: []resource.TestStep{
// 			{
// 				Config:      testAccDataprocJob_missingJobConf(),
// 				ExpectError: regexp.MustCompile("You must define and configure exactly one xxx_config block"),
// 			},
// 		},
// 	})
// }

func TestAccDataprocJob_updatable(t *testing.T) {
	t.Parallel()

	var job dataproc.Job
	rnd := randString(t, 10)
	jobId := fmt.Sprintf("dproc-update-job-id-%s", rnd)
	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDataprocJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocJob_updatable(rnd, jobId, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocJobExists(t, "google_dataproc_job.updatable", &job),
					resource.TestCheckResourceAttr("google_dataproc_job.updatable", "force_delete", "false"),
				),
			},
			{
				Config: testAccDataprocJob_updatable(rnd, jobId, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocJobExists(t, "google_dataproc_job.updatable", &job),
					resource.TestCheckResourceAttr("google_dataproc_job.updatable", "force_delete", "true"),
				),
			},
		},
	})
}

func TestAccDataprocJob_PySpark(t *testing.T) {
	t.Parallel()

	var job dataproc.Job
	rnd := randString(t, 10)
	jobId := fmt.Sprintf("dproc-custom-job-id-%s", rnd)
	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDataprocJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocJob_pySpark(rnd),
				Check: resource.ComposeTestCheckFunc(

					testAccCheckDataprocJobExists(t, "google_dataproc_job.pyspark", &job),

					// Custom supplied job_id
					resource.TestCheckResourceAttr("google_dataproc_job.pyspark", "reference.0.job_id", jobId),

					// Autogenerated / computed values
					resource.TestCheckResourceAttrSet("google_dataproc_job.pyspark", "status.0.state"),
					resource.TestCheckResourceAttrSet("google_dataproc_job.pyspark", "status.0.state_start_time"),
					resource.TestCheckResourceAttr("google_dataproc_job.pyspark", "scheduling.0.max_failures_per_hour", "1"),
					resource.TestCheckResourceAttr("google_dataproc_job.pyspark", "scheduling.0.max_failures_total", "20"),
					resource.TestCheckResourceAttr("google_dataproc_job.pyspark", "labels.one", "1"),

					// Unique job config
					testAccCheckDataprocJobAttrMatch(
						"google_dataproc_job.pyspark", "pyspark_config", &job),

					// Wait until job completes successfully
					testAccCheckDataprocJobCompletesSuccessfully(t, "google_dataproc_job.pyspark", &job),
				),
			},
		},
	})
}

func TestAccDataprocJob_Spark(t *testing.T) {
	t.Parallel()

	var job dataproc.Job
	rnd := randString(t, 10)
	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDataprocJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocJob_spark(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocJobExists(t, "google_dataproc_job.spark", &job),

					// Autogenerated / computed values
					resource.TestCheckResourceAttrSet("google_dataproc_job.spark", "reference.0.job_id"),
					resource.TestCheckResourceAttrSet("google_dataproc_job.spark", "status.0.state"),
					resource.TestCheckResourceAttrSet("google_dataproc_job.spark", "status.0.state_start_time"),

					// Unique job config
					testAccCheckDataprocJobAttrMatch(
						"google_dataproc_job.spark", "spark_config", &job),

					// Wait until job completes successfully
					testAccCheckDataprocJobCompletesSuccessfully(t, "google_dataproc_job.spark", &job),
				),
			},
		},
	})
}

func TestAccDataprocJob_Hadoop(t *testing.T) {
	t.Parallel()

	var job dataproc.Job
	rnd := randString(t, 10)
	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDataprocJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocJob_hadoop(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocJobExists(t, "google_dataproc_job.hadoop", &job),

					// Autogenerated / computed values
					resource.TestCheckResourceAttrSet("google_dataproc_job.hadoop", "reference.0.job_id"),
					resource.TestCheckResourceAttrSet("google_dataproc_job.hadoop", "status.0.state"),
					resource.TestCheckResourceAttrSet("google_dataproc_job.hadoop", "status.0.state_start_time"),

					// Unique job config
					testAccCheckDataprocJobAttrMatch(
						"google_dataproc_job.hadoop", "hadoop_config", &job),

					// Wait until job completes successfully
					testAccCheckDataprocJobCompletesSuccessfully(t, "google_dataproc_job.hadoop", &job),
				),
			},
		},
	})
}

func TestAccDataprocJob_Hive(t *testing.T) {
	t.Parallel()

	var job dataproc.Job
	rnd := randString(t, 10)
	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDataprocJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocJob_hive(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocJobExists(t, "google_dataproc_job.hive", &job),

					// Autogenerated / computed values
					resource.TestCheckResourceAttrSet("google_dataproc_job.hive", "reference.0.job_id"),
					resource.TestCheckResourceAttrSet("google_dataproc_job.hive", "status.0.state"),
					resource.TestCheckResourceAttrSet("google_dataproc_job.hive", "status.0.state_start_time"),

					// Unique job config
					testAccCheckDataprocJobAttrMatch(
						"google_dataproc_job.hive", "hive_config", &job),

					// Wait until job completes successfully
					testAccCheckDataprocJobCompletesSuccessfully(t, "google_dataproc_job.hive", &job),
				),
			},
		},
	})
}

func TestAccDataprocJob_Pig(t *testing.T) {
	t.Parallel()

	var job dataproc.Job
	rnd := randString(t, 10)
	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDataprocJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocJob_pig(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocJobExists(t, "google_dataproc_job.pig", &job),

					// Autogenerated / computed values
					resource.TestCheckResourceAttrSet("google_dataproc_job.pig", "reference.0.job_id"),
					resource.TestCheckResourceAttrSet("google_dataproc_job.pig", "status.0.state"),
					resource.TestCheckResourceAttrSet("google_dataproc_job.pig", "status.0.state_start_time"),

					// Unique job config
					testAccCheckDataprocJobAttrMatch(
						"google_dataproc_job.pig", "pig_config", &job),

					// Wait until job completes successfully
					testAccCheckDataprocJobCompletesSuccessfully(t, "google_dataproc_job.pig", &job),
				),
			},
		},
	})
}

func TestAccDataprocJob_SparkSql(t *testing.T) {
	t.Parallel()

	var job dataproc.Job
	rnd := randString(t, 10)
	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDataprocJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocJob_sparksql(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocJobExists(t, "google_dataproc_job.sparksql", &job),

					// Autogenerated / computed values
					resource.TestCheckResourceAttrSet("google_dataproc_job.sparksql", "reference.0.job_id"),
					resource.TestCheckResourceAttrSet("google_dataproc_job.sparksql", "status.0.state"),
					resource.TestCheckResourceAttrSet("google_dataproc_job.sparksql", "status.0.state_start_time"),

					// Unique job config
					testAccCheckDataprocJobAttrMatch(
						"google_dataproc_job.sparksql", "sparksql_config", &job),

					// Wait until job completes successfully
					testAccCheckDataprocJobCompletesSuccessfully(t, "google_dataproc_job.sparksql", &job),
				),
			},
		},
	})
}

func TestAccDataprocJob_Presto(t *testing.T) {
	t.Parallel()

	var job dataproc.Job
	rnd := randString(t, 10)
	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDataprocJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocJob_presto(rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocJobExists(t, "google_dataproc_job.presto", &job),

					// Autogenerated / computed values
					resource.TestCheckResourceAttrSet("google_dataproc_job.presto", "reference.0.job_id"),
					resource.TestCheckResourceAttrSet("google_dataproc_job.presto", "status.0.state"),
					resource.TestCheckResourceAttrSet("google_dataproc_job.presto", "status.0.state_start_time"),

					// Unique job config
					testAccCheckDataprocJobAttrMatch(
						"google_dataproc_job.presto", "presto_config", &job),

					// Wait until job completes successfully
					testAccCheckDataprocJobCompletesSuccessfully(t, "google_dataproc_job.presto", &job),
				),
			},
		},
	})
}

func testAccCheckDataprocJobDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		config := googleProviderConfig(t)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "google_dataproc_job" {
				continue
			}

			if rs.Primary.ID == "" {
				return fmt.Errorf("Unable to verify delete of dataproc job ID is empty")
			}
			attributes := rs.Primary.Attributes

			project, err := getTestProject(rs.Primary, config)
			if err != nil {
				return err
			}

			parts := strings.Split(rs.Primary.ID, "/")
			job_id := parts[len(parts)-1]
			_, err = config.NewDataprocClient(config.userAgent).Projects.Regions.Jobs.Get(
				project, attributes["region"], job_id).Do()
			if err != nil {
				if gerr, ok := err.(*googleapi.Error); ok && gerr.Code == 404 {
					return nil
				} else if ok {
					return fmt.Errorf("Error making GCP platform call: http code error : %d, http message error: %s", gerr.Code, gerr.Message)
				}
				return fmt.Errorf("Error making GCP platform call: %s", err.Error())
			}
			return fmt.Errorf("Dataproc job still exists")
		}

		return nil
	}
}

func testAccCheckDataprocJobCompletesSuccessfully(t *testing.T, n string, job *dataproc.Job) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := googleProviderConfig(t)

		attributes := s.RootModule().Resources[n].Primary.Attributes
		region := attributes["region"]
		project, err := getTestProject(s.RootModule().Resources[n].Primary, config)
		if err != nil {
			return err
		}

		jobCompleteTimeoutMins := 5 * time.Minute
		waitErr := dataprocJobOperationWait(config, region, project, job.Reference.JobId,
			"Awaiting Dataproc job completion", config.userAgent, jobCompleteTimeoutMins)
		if waitErr != nil {
			return waitErr
		}

		completeJob, err := config.NewDataprocClient(config.userAgent).Projects.Regions.Jobs.Get(
			project, region, job.Reference.JobId).Do()
		if err != nil {
			return err
		}
		if completeJob.Status.State == "ERROR" {
			if !strings.HasPrefix(completeJob.DriverOutputResourceUri, "gs://") {
				return fmt.Errorf("Job completed in ERROR state but no valid log URI found")
			}
			u := strings.SplitN(strings.TrimPrefix(completeJob.DriverOutputResourceUri, "gs://"), "/", 2)
			if len(u) != 2 {
				return fmt.Errorf("Job completed in ERROR state but no valid log URI found")
			}
			l, err := config.NewStorageClient(config.userAgent).Objects.List(u[0]).Prefix(u[1]).Do()
			if err != nil {
				return errwrap.Wrapf("Job completed in ERROR state, found error when trying to list logs: {{err}}", err)
			}
			for _, item := range l.Items {
				resp, err := config.NewStorageClient(config.userAgent).Objects.Get(item.Bucket, item.Name).Download()
				if err != nil {
					return errwrap.Wrapf("Job completed in ERROR state, found error when trying to read logs: {{err}}", err)
				}
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return errwrap.Wrapf("Job completed in ERROR state, found error when trying to read logs: {{err}}", err)
				}
				log.Printf("[ERROR] Job failed, driver logs:\n%s", body)
			}
			return fmt.Errorf("Job completed in ERROR state, check logs for details")
		} else if completeJob.Status.State != "DONE" && completeJob.Status.State != "RUNNING" {
			return fmt.Errorf("Job did not complete successfully, instead status: %s", completeJob.Status.State)
		}

		return nil
	}
}

func testAccCheckDataprocJobExists(t *testing.T, n string, job *dataproc.Job) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Terraform resource Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set for Dataproc job")
		}

		config := googleProviderConfig(t)
		parts := strings.Split(s.RootModule().Resources[n].Primary.ID, "/")
		jobId := parts[len(parts)-1]
		project, err := getTestProject(s.RootModule().Resources[n].Primary, config)
		if err != nil {
			return err
		}

		found, err := config.NewDataprocClient(config.userAgent).Projects.Regions.Jobs.Get(
			project, rs.Primary.Attributes["region"], jobId).Do()
		if err != nil {
			return err
		}

		*job = *found

		return nil
	}
}

func testAccCheckDataprocJobAttrMatch(n, jobType string, job *dataproc.Job) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		attributes, err := getResourceAttributes(n, s)
		if err != nil {
			return err
		}

		jobTests := []jobTestField{}
		if jobType == "pyspark_config" {
			jobTests = append(jobTests, jobTestField{"pyspark_config.0.main_python_file_uri", job.PysparkJob.MainPythonFileUri})
			jobTests = append(jobTests, jobTestField{"pyspark_config.0.args", job.PysparkJob.Args})
			jobTests = append(jobTests, jobTestField{"pyspark_config.0.python_file_uris", job.PysparkJob.PythonFileUris})
			jobTests = append(jobTests, jobTestField{"pyspark_config.0.jar_file_uris", job.PysparkJob.JarFileUris})
			jobTests = append(jobTests, jobTestField{"pyspark_config.0.file_uris", job.PysparkJob.FileUris})
			jobTests = append(jobTests, jobTestField{"pyspark_config.0.archive_uris", job.PysparkJob.ArchiveUris})
			jobTests = append(jobTests, jobTestField{"pyspark_config.0.properties", job.PysparkJob.Properties})
			jobTests = append(jobTests, jobTestField{"pyspark_config.0.logging_config.0.driver_log_levels", job.PysparkJob.LoggingConfig.DriverLogLevels})
		}
		if jobType == "spark_config" {
			jobTests = append(jobTests, jobTestField{"spark_config.0.main_class", job.SparkJob.MainClass})
			jobTests = append(jobTests, jobTestField{"spark_config.0.main_jar_file_uri", job.SparkJob.MainJarFileUri})
			jobTests = append(jobTests, jobTestField{"spark_config.0.args", job.SparkJob.Args})
			jobTests = append(jobTests, jobTestField{"spark_config.0.jar_file_uris", job.SparkJob.JarFileUris})
			jobTests = append(jobTests, jobTestField{"spark_config.0.file_uris", job.SparkJob.FileUris})
			jobTests = append(jobTests, jobTestField{"spark_config.0.archive_uris", job.SparkJob.ArchiveUris})
			jobTests = append(jobTests, jobTestField{"spark_config.0.properties", job.SparkJob.Properties})
			jobTests = append(jobTests, jobTestField{"spark_config.0.logging_config.0.driver_log_levels", job.SparkJob.LoggingConfig.DriverLogLevels})
		}
		if jobType == "hadoop_config" {
			jobTests = append(jobTests, jobTestField{"hadoop_config.0.main_class", job.HadoopJob.MainClass})
			jobTests = append(jobTests, jobTestField{"hadoop_config.0.main_jar_file_uri", job.HadoopJob.MainJarFileUri})
			jobTests = append(jobTests, jobTestField{"hadoop_config.0.args", job.HadoopJob.Args})
			jobTests = append(jobTests, jobTestField{"hadoop_config.0.jar_file_uris", job.HadoopJob.JarFileUris})
			jobTests = append(jobTests, jobTestField{"hadoop_config.0.file_uris", job.HadoopJob.FileUris})
			jobTests = append(jobTests, jobTestField{"hadoop_config.0.archive_uris", job.HadoopJob.ArchiveUris})
			jobTests = append(jobTests, jobTestField{"hadoop_config.0.properties", job.HadoopJob.Properties})
			jobTests = append(jobTests, jobTestField{"hadoop_config.0.logging_config.0.driver_log_levels", job.HadoopJob.LoggingConfig.DriverLogLevels})
		}
		if jobType == "hive_config" {
			queries := []string{}
			if job.HiveJob.QueryList != nil {
				queries = job.HiveJob.QueryList.Queries
			}
			jobTests = append(jobTests, jobTestField{"hive_config.0.query_list", queries})
			jobTests = append(jobTests, jobTestField{"hive_config.0.query_file_uri", job.HiveJob.QueryFileUri})
			jobTests = append(jobTests, jobTestField{"hive_config.0.continue_on_failure", job.HiveJob.ContinueOnFailure})
			jobTests = append(jobTests, jobTestField{"hive_config.0.script_variables", job.HiveJob.ScriptVariables})
			jobTests = append(jobTests, jobTestField{"hive_config.0.properties", job.HiveJob.Properties})
			jobTests = append(jobTests, jobTestField{"hive_config.0.jar_file_uris", job.HiveJob.JarFileUris})
		}
		if jobType == "pig_config" {
			queries := []string{}
			if job.PigJob.QueryList != nil {
				queries = job.PigJob.QueryList.Queries
			}
			jobTests = append(jobTests, jobTestField{"pig_config.0.query_list", queries})
			jobTests = append(jobTests, jobTestField{"pig_config.0.query_file_uri", job.PigJob.QueryFileUri})
			jobTests = append(jobTests, jobTestField{"pig_config.0.continue_on_failure", job.PigJob.ContinueOnFailure})
			jobTests = append(jobTests, jobTestField{"pig_config.0.script_variables", job.PigJob.ScriptVariables})
			jobTests = append(jobTests, jobTestField{"pig_config.0.properties", job.PigJob.Properties})
			jobTests = append(jobTests, jobTestField{"pig_config.0.jar_file_uris", job.PigJob.JarFileUris})
		}
		if jobType == "sparksql_config" {
			queries := []string{}
			if job.SparkSqlJob.QueryList != nil {
				queries = job.SparkSqlJob.QueryList.Queries
			}
			jobTests = append(jobTests, jobTestField{"sparksql_config.0.query_list", queries})
			jobTests = append(jobTests, jobTestField{"sparksql_config.0.query_file_uri", job.SparkSqlJob.QueryFileUri})
			jobTests = append(jobTests, jobTestField{"sparksql_config.0.script_variables", job.SparkSqlJob.ScriptVariables})
			jobTests = append(jobTests, jobTestField{"sparksql_config.0.properties", job.SparkSqlJob.Properties})
			jobTests = append(jobTests, jobTestField{"sparksql_config.0.jar_file_uris", job.SparkSqlJob.JarFileUris})
		}

		for _, attrs := range jobTests {
			if c := checkMatch(attributes, attrs.tf_attr, attrs.gcp_attr); c != "" {
				return fmt.Errorf(c)
			}
		}

		return nil
	}
}

// TODO (mbang): Test `ExactlyOneOf` here
// func testAccDataprocJob_missingJobConf() string {
// 	return `
// resource "google_dataproc_job" "missing_config" {
// 	placement {
// 		cluster_name = "na"
// 	}

// 	force_delete = true
// }`
// }

var singleNodeClusterConfig = `
resource "google_dataproc_cluster" "basic" {
  name   = "dproc-job-test-%s"
  region = "us-central1"

  cluster_config {
    # Keep the costs down with smallest config we can get away with
    software_config {
      override_properties = {
        "dataproc:dataproc.allow.zero.workers" = "true"
      }
    }

    master_config {
      num_instances = 1
      machine_type  = "e2-standard-2"
      disk_config {
        boot_disk_size_gb = 35
      }
    }
  }
}
`

func testAccDataprocJob_updatable(rnd, jobId, del string) string {
	return fmt.Sprintf(
		singleNodeClusterConfig+`
resource "google_dataproc_job" "updatable" {
  placement {
    cluster_name = google_dataproc_cluster.basic.name
  }
  reference {
    job_id = "%s"
  }

  region       = google_dataproc_cluster.basic.region
  force_delete = %s

  pyspark_config {
    main_python_file_uri = "gs://dataproc-examples-2f10d78d114f6aaec76462e3c310f31f/src/pyspark/hello-world/hello-world.py"
  }
}
`, rnd, jobId, del)
}

func testAccDataprocJob_pySpark(rnd string) string {
	return fmt.Sprintf(
		singleNodeClusterConfig+`
resource "google_dataproc_job" "pyspark" {
  placement {
    cluster_name = google_dataproc_cluster.basic.name
  }
  reference {
    job_id = "dproc-custom-job-id-%s"
  }

  region       = google_dataproc_cluster.basic.region
  force_delete = true

  pyspark_config {
    main_python_file_uri = "gs://dataproc-examples-2f10d78d114f6aaec76462e3c310f31f/src/pyspark/hello-world/hello-world.py"
    properties = {
      "spark.logConf" = "true"
    }
    logging_config {
      driver_log_levels = {
        "root" = "INFO"
      }
    }
  }

  scheduling {
	max_failures_per_hour = 1
	max_failures_total=20
  }

  labels = {
    one = "1"
  }
}
`, rnd, rnd)
}

func testAccDataprocJob_spark(rnd string) string {
	return fmt.Sprintf(
		singleNodeClusterConfig+`
resource "google_dataproc_job" "spark" {
  region       = google_dataproc_cluster.basic.region
  force_delete = true
  placement {
    cluster_name = google_dataproc_cluster.basic.name
  }

  spark_config {
    main_class    = "org.apache.spark.examples.SparkPi"
    jar_file_uris = ["file:///usr/lib/spark/examples/jars/spark-examples.jar"]
    args          = ["1000"]
    properties = {
      "spark.logConf" = "true"
    }
    logging_config {
      driver_log_levels = {
      }
    }
  }
}
`, rnd)

}

func testAccDataprocJob_hadoop(rnd string) string {
	return fmt.Sprintf(
		singleNodeClusterConfig+`
resource "google_dataproc_job" "hadoop" {
  region       = google_dataproc_cluster.basic.region
  force_delete = true
  placement {
    cluster_name = google_dataproc_cluster.basic.name
  }

  hadoop_config {
    main_jar_file_uri = "file:///usr/lib/hadoop-mapreduce/hadoop-mapreduce-examples.jar"
    args = [
      "wordcount",
      "file:///usr/lib/spark/NOTICE",
      "gs://${google_dataproc_cluster.basic.cluster_config[0].bucket}/hadoopjob_output_%s",
    ]
  }
}
`, rnd, rnd)

}

func testAccDataprocJob_hive(rnd string) string {
	return fmt.Sprintf(
		singleNodeClusterConfig+`
resource "google_dataproc_job" "hive" {
  region       = google_dataproc_cluster.basic.region
  force_delete = true
  placement {
    cluster_name = google_dataproc_cluster.basic.name
  }

  hive_config {
    query_list = [
      "DROP TABLE IF EXISTS dprocjob_test",
      "CREATE EXTERNAL TABLE dprocjob_test(bar int) LOCATION 'gs://${google_dataproc_cluster.basic.cluster_config[0].bucket}/hive_dprocjob_test/'",
      "SELECT * FROM dprocjob_test WHERE bar > 2",
    ]
  }
}
`, rnd)

}

func testAccDataprocJob_pig(rnd string) string {
	return fmt.Sprintf(
		singleNodeClusterConfig+`
resource "google_dataproc_job" "pig" {
  region       = google_dataproc_cluster.basic.region
  force_delete = true
  placement {
    cluster_name = google_dataproc_cluster.basic.name
  }

  pig_config {
    query_list = [
      "LNS = LOAD 'file:///usr/lib/pig/LICENSE.txt ' AS (line)",
      "WORDS = FOREACH LNS GENERATE FLATTEN(TOKENIZE(line)) AS word",
      "GROUPS = GROUP WORDS BY word",
      "WORD_COUNTS = FOREACH GROUPS GENERATE group, COUNT(WORDS)",
      "DUMP WORD_COUNTS",
    ]
  }
}
`, rnd)

}

func testAccDataprocJob_sparksql(rnd string) string {
	return fmt.Sprintf(
		singleNodeClusterConfig+`
resource "google_dataproc_job" "sparksql" {
  region       = google_dataproc_cluster.basic.region
  force_delete = true
  placement {
    cluster_name = google_dataproc_cluster.basic.name
  }

  sparksql_config {
    query_list = [
      "DROP TABLE IF EXISTS dprocjob_test",
      "CREATE TABLE dprocjob_test(bar int)",
      "SELECT * FROM dprocjob_test WHERE bar > 2",
    ]
  }
}
`, rnd)

}

func testAccDataprocJob_presto(rnd string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "basic" {
  name   = "dproc-job-test-%s"
  region = "us-central1"

  cluster_config {
    # Keep the costs down with smallest config we can get away with
    software_config {
      override_properties = {
        "dataproc:dataproc.allow.zero.workers" = "true"
      }
	  optional_components = ["PRESTO"]
    }

    master_config {
      num_instances = 1
      machine_type  = "e2-standard-2"
      disk_config {
        boot_disk_size_gb = 35
      }
    }
  }
}

resource "google_dataproc_job" "presto" {
  region       = google_dataproc_cluster.basic.region
  force_delete = true
  placement {
    cluster_name = google_dataproc_cluster.basic.name
  }

  presto_config {
    query_list = [
      "SELECT * FROM system.metadata.schema_properties"
    ]
  }
}
`, rnd)

}
