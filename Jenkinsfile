pipeline {
    agent any

    environment {
        AWS_ACCESS_KEY_ID = credentials('AWS_ACCESS_KEY_ID')
        AWS_SECRET_ACCESS_KEY = credentials('AWS_SECRET_ACCESS_KEY')
    }

    stages {
        stage('Checkout') {
            steps {
                dir("app") {
                    git branch: 'main', url: 'https://github.com/jiechenmc/Freon.git'
                }
            }
        }
        stage('Terraform') {
            steps {
                script {
                        dir('app') {
                            sh 'terraform init'
                            sh 'terraform validate'
                            sh "terraform apply -auto-approve"
                            }
                        }
                }
        }
    }
}