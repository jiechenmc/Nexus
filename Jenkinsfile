pipeline {
    agent any

    parameters {
        choice(name: 'ACTION', choices: ['apply', 'destroy'], description: 'What action should Terraform take?')
    }

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
        stage('List Directory') {
            steps {
                sh 'ls **'
            }
        }
        stage('Terraform') {
            steps {
                script {
                        dir('app') {
                            sh 'terraform init'
                            sh 'terraform validate'
                            sh "terraform ${params.ACTION} -auto-approve"
                            }
                        }
                }
        }
    }
}