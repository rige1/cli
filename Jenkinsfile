pipeline {
    agent {
        label "amd64 && ubuntu-1804 && overlay2"
    }

    options {
        timeout(time: 60, unit: 'MINUTES')
    }

    stages {
        stage("Docker info") {
            steps {
                sh "docker version"
                sh "docker info"
            }
        }
        stage("e2e (non-experimental) - stable engine") {
            steps {
                sh "E2E_UNIQUE_ID=clie2e${BUILD_NUMBER} \
                    IMAGE_TAG=clie2e${BUILD_NUMBER} \
                    make -f docker.Makefile test-e2e-non-experimental"
            }
        }
        stage("e2e (non-experimental) - 19.03 engine") {
            steps {
                sh "E2E_ENGINE_VERSION=19.03-dind \
                  E2E_UNIQUE_ID=clie2e${BUILD_NUMBER} \
                  IMAGE_TAG=clie2e${BUILD_NUMBER} \
                  make -f docker.Makefile test-e2e-non-experimental"
            }
        }
        stage("e2e (experimental)") {
            steps {
                sh "E2E_UNIQUE_ID=clie2e${BUILD_NUMBER} \
                    IMAGE_TAG=clie2e${BUILD_NUMBER} \
                    make -f docker.Makefile test-e2e-experimental"
            }
        }
        stage("e2e (ssh connhelper)") {
            steps {
                sh "E2E_UNIQUE_ID=clie2e${BUILD_NUMBER} \
                    IMAGE_TAG=clie2e${BUILD_NUMBER} \
                    make -f docker.Makefile test-e2e-connhelper-ssh"
            }
        }
    }
}
