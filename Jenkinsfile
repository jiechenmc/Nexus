pipeline {
    agent any

    environment {
        AWS_ACCESS_KEY_ID = credentials('AWS_ACCESS_KEY_ID')
        AWS_SECRET_ACCESS_KEY = credentials('AWS_SECRET_ACCESS_KEY')
    }

    stages {
        stage('Checkout') {
            steps {
             checkout scmGit(
                branches: [[name: 'main']],
                userRemoteConfigs: [[url: 'https://github.com/jiechenmc/Freon.git']])
            }
        }
        stage('Terraform') {
            steps {
                script {
                        dir('Freon') {
                            sh 'terraform init'
                            sh 'terraform validate'
                            sh "terraform apply -auto-approve"
                            }
                        }
                }
        }
    }
}