pipeline {
    agent any

    parameters {
        choice(name: 'ACTION', choices: ['apply', 'destroy'], description: 'What action should Terraform take?')
    }

    environment {
        AWS_ACCESS_KEY_ID = credentials('AWS_ACCESS_KEY_ID')
        AWS_SECRET_ACCESS_KEY = credentials('AWS_SECRET_ACCESS_KEY')
        ANSIBLE_HOST_KEY_CHECKING = "False"
    }

    stages {
        stage('List Directory') {
            steps {
                sh 'ls **'
            }
        }
        stage('Terraform') {
            steps {
                script {
                        dir('terraform') {
                            sh 'terraform init'
                            sh 'terraform validate'
                            sh "terraform ${params.ACTION} -auto-approve"
                            }
                        }
                sleep(time:10,unit:"SECONDS")
                }
        }
        stage('Ansible'){
            when {
                expression {
                    return params.ACTION == 'apply'
                }
            }
            steps {
                withCredentials([sshUserPrivateKey(credentialsId: 'wsl', keyFileVariable: 'SSH_KEYFILE', usernameVariable: 'ubuntu')]) {
                    sh 'ansible-playbook ./ansible/site.yml -i hosts -u ubuntu --private-key=${SSH_KEYFILE}' 
                }
            }
        }
    }
}