# Jengo
## How to Install it
### MacOS
```
brew tap tkennes/jengo

brew install jengo
```


## How To Use it

### Step 0. - Run a local Jenkins Server
If you already have a Jenkins Server, continue with the next step. If you are looking to deploy a more production-worthy Jenkins server, consult the Jenkins documentation.

Shortest option:

```bash
docker pull jenkins/jenkins

docker run -p 8080:8080 -p 50000:50000 -v jenkins_home:/var/jenkins_home jenkins/jenkins:lts
```

Short option: https://medium.com/@YetkinTimocin/creating-a-local-jenkins-server-using-docker-2e4dfe7b5880


Long option: https://www.jenkins.io/doc/book/installing/docker/

### Step 1. Create an API Token
If you have sufficient rights, you should be able to create a token at the url suffix: /user/<username>/configure.

See also the screenshot below:
<img src="./static/create_token.png">

### Step 1. Create a configuraiton file
Name it: 
```
~/.jenkins.yaml
```
Fill it:
```yaml
current: localhost
contexts: 
 - name: localhost
   url: "http://localhost:8080"
   username: "<username>"
   token:    "<token>"
```

In the future, I might add options for more customization with regard to the contexts, as well as making use of multiple contexts. Currently, this is not of direct use for my own projects, so I am postponing this. 

Contributions are as always much appreciated!

#### Step 2. 
The commands should be working now!

Give it a try:
```
jengo jobs
```

```
jengo builds --job <job-name>
```

## Contribute
We make use of githooks, taken out of the unsynced ./.git/hooks folder and put in the .githooks folder (surprisingly :)). This allows for versioning of githooks as well. There is only a small caveat, you need to adjust your git settings before it is able to pick up this change. If the git your are using is of version 2.29 or above, doing so is fairly simple:
```
git config core.hooksPath .githooks
```

Notice that this only applies to the repository you're currently in.

https://www.viget.com/articles/two-ways-to-share-git-hooks-with-your-team/