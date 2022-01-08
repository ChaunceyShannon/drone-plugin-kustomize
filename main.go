package main

import (
	. "github.com/ChaunceyShannon/golanglibs"
)

func system(cmd string) {
	if Os.System(cmd) != 0 {
		Panicerr("Error while execute command")
	}
}

func main() {
	git_username := Os.Getenv("PLUGIN_GIT_USERNAME")
	git_password := Os.Getenv("PLUGIN_GIT_PASSWORD")
	git_repo := Tools.URL(Os.Getenv("PLUGIN_GIT_REPO")).Parse()
	git_branch := Os.Getenv("PLUGIN_GIT_BRANCH")
	git_app_path := Os.Getenv("PLUGIN_GIT_APP_PATH")
	docker_image := Os.Getenv("PLUGIN_DOCKER_IMAGE") + ":" + Os.Getenv("PLUGIN_DOCKER_TAG")

	tmpGitDir := "/tmp/git_repo"
	Os.Mkdir(tmpGitDir)
	Os.Chdir(tmpGitDir)

	Lg.Trace("Using temporary directory:", tmpGitDir)

	Lg.Trace("Clone git repository")
	system("git clone -b " + git_branch + " " + git_repo.Schema + "://" + git_username + ":" + git_password + "@" + git_repo.Host + git_repo.Path + " .")

	Lg.Trace("Change directory to", git_app_path)
	Os.Mkdir(Os.Path.Join(tmpGitDir, git_app_path))
	Os.Chdir(Os.Path.Join(tmpGitDir, git_app_path))

	Lg.Trace("Set docker image with kustomize:", docker_image)
	if !Os.Path.Exists("kustomization.yaml") {
		Os.Touch("kustomization.yaml")
	}
	system("kustomize edit set image " + docker_image)

	Os.Chdir(tmpGitDir)

	Lg.Trace("Push to git repository")
	system("git config --global user.email 'drone@" + git_repo.Host + "'")
	system("git config --global user.name 'drone'")
	system("git add *")
	system("git commit -m 'Set image to " + docker_image)
	system("git push")
}
