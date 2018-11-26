#!groovy
@Library('blockchyp-pipelines') _

pipeline {
  agent none

  stages {
    stage('Start') {
      agent {
        node { label 'docker' }
      }
      steps {
        notifySlack 'Started'
      }
    }

    stage('Build Go') {
      agent {
        docker {
          registryUrl 'https://105741232974.dkr.ecr.us-west-2.amazonaws.com'
          registryCredentialsId 'ecr:us-west-2:f4cba16d-79b0-4e7e-b771-921bf1ed54fa'
          image 'blockchyp/go-build'
          args '-v /tmp/go/pkg/mod:/home/build/go/pkg/mod'
        }
      }

      environment {
        GITHUB = credentials('github')
      }

      stages {
        stage('Lint') {
          steps {
            sh "make lint"
          }
        }

        stage('Test') {
          steps {
            sh "make test"
          }
        }
      }

      post {
        always {
          junit allowEmptyResults: true, testResults: 'build/test-reports/*.xml'
        }
      }
    }

    stage('SonarQube') {
      agent {
        node { label 'docker' }
      }

      steps {
        withSonarQubeEnv('SonarQubeDev') {
          sh "${tool 'SonarQube Scanner 3.2.0.1227'}/bin/sonar-scanner"
        }
      }
    }
  }

  post {
    always {
      notifySlack()
    }
  }
}
