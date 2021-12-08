## New Release Guide for VDO



### Pre-requisite
1. Check the status of all the Open Issues, cross-check if any of it is critical and must be fixed before the release.
2. Check for the Open PR's and cross-check if all the important and Must-Fix PR's are merged.
3. Visit the latest Milestone and see if anything important issues needs to be addressed.
4. Check the status of CI (Buildl & Deploy) and see if con-current builds have passed.

### Make a new Tag
Before making a new Tag you must clone and build & deploy the code locally to verify the basic features of the project.
```
cd $GOPATH/src
mkdir -p github.com/vmware-tanzu
cd github.com/vmware-tanzu
git clone git@github.com:vmware-tanzu/vsphere-kubernetes-drivers-operator.git
```

Create a new Tag

```
cd $GOPATH/src/github.com/vmware-tanzu/vsphere-kubernetes-drivers-operator

git add tag -a 0.1.1 -m "Release 0.1.1"
#Replace 0.1.1 with your new version number
```

Checkout the Code and run deploy command 

```
git checkout 0.1.1

make deploy
```

If the Deploy command passes then your DockerImage would have been made successfully

```
docker images | grep vmware.com/vdo

```

Copy the latest TAG from the `vmware.com/vdo` image and make a new tag

```

```
