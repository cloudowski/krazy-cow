pipeline {
    // In multibranch it's not necessary
    // triggers {
    //     pollSCM("")
    // }
    agent none
    
    stages { 
        stage('Build app') {
            agent {
                kubernetes {
                    containerTemplate {
                        name 'kubectl'
                        image 'golang:1.12-buster'
                        ttyEnabled true
                        command 'cat'
                    }
                }
            }
            steps {
                // sh 'env|sort'
                sh 'go get -d'
                sh 'make test'
                sh 'make build'
                // sh 'echo -e "#!/bin/sh\ncat" > cow;chmod +x cow'
                stash name: "app", includes: "cow"
            }
        }
        stage('Build image') {
            agent {
                kubernetes {
                    yamlFile "ci/kaniko.yaml"
                }
            }
            steps {
                container(name: 'kaniko', shell: '/busybox/sh') {
                    unstash 'app'
                    sh '''
                    /busybox/sh ci/getversion.sh > .ci_version
                    ver=`cat .ci_version`
                    ci/build-kaniko.sh cloudowski/test1:\$ver Dockerfile.embed
                    '''
                }
            }
            // when {
                //  commit message starts with build
                // changelog '^build .+$'
            // }
        }

        stage('Deploy preview') {
            agent {
                kubernetes {
                    serviceAccount 'deployer'
                    containerTemplate {
                        name 'kubectl'
                        image 'cloudowski/drone-kustomize'
                        ttyEnabled true
                        command 'cat'
                    }
                }
            }
            steps {
                sh 'ci/deploy-kustomize.sh -p'
                // rocketSend channel: 'general', message: "Visit me @ $BUILD_URL"
            }
            when {
                // deploy on PR automatically OR on non-mastyer when commit starts with "shipit"
                anyOf {
                    allOf {
                        changelog '^shipit ?.+$'
                        expression {
                            ! (env.GIT_BRANCH =~ /^(master|testmaster)$/)
                        }
                    }
                    expression {
                        (env.BRANCH_NAME =~ /^PR-.*/)
                    }
                }
            }
        }

        stage('Deploy stage') {
            agent {
                kubernetes {
                    serviceAccount 'deployer'
                    containerTemplate {
                        name 'kubectl'
                        image 'cloudowski/drone-kustomize'
                        ttyEnabled true
                        command 'cat'
                    }
                }
            }
            steps {
                sh 'ci/deploy-kustomize.sh -t kcow-stage'
                // rocketSend channel: 'general', message: "Visit me @ $BUILD_URL"
            }
            when {
                allOf {
                    // changelog '^deploy ?.+$'
                    expression {
                        (env.GIT_BRANCH =~ /^(master|testmaster)$/)
                    }
                }
            }
        }
        

        stage('Deploy prod') {
            agent {
                kubernetes {
                    serviceAccount 'deployer'
                    containerTemplate {
                        name 'kubectl'
                        image 'cloudowski/drone-kustomize'
                        ttyEnabled true
                        command 'cat'
                    }
                }
            }

            input {
                message "Deploy to prod?"
                ok "Ship it!"
            }

            steps {
                sh 'ci/deploy-kustomize.sh -t kcow-prod'
                // rocketSend channel: 'general', message: "Visit me @ $BUILD_URL"
            }
            when {
                // beforeInput true
                expression {
                    (env.GIT_BRANCH =~ /^(master|testmaster)$/)
                }
            }
        }
    }
}
