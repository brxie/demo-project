
node('demo') {
    env.GOPATH = "${env.WORKSPACE}"
    env.GOOS = 'linux'
    env.PATH   = "${env.PATH}:/usr/local/go/bin/:/go/bin"

    stage('Checkout') {
        deleteDir()
        checkout scm
    }

    stage('Build') {
        sh 'go build -o app'
    }

    stage('Unit test') {
        sh 'go test ./... -v | go-junit-report > report.xml'
        junit 'report.xml'
    }

    stage('Generate docs') {
        sh 'godoc cmd/base64/ > docs.txt'
    }

    stage('Archive artifacts') {
        archiveArtifacts 'app'
        archiveArtifacts 'docs.txt'
    }
}

