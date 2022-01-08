# drone-plugin-kustomize

Usage 

```yaml
- name: update to k8s
  image: chaunceyshannon/kustomize:v4.4.1
  settings:
    git_username: user
    git_password: pass
    git_repo: https://gitea.example.com/gitea/flux
    git_branch: master
    git_app_path: /app/test
    docker_image: registry.example.com/demo
    docker_tag: ${DRONE_COMMIT_SHA}
  when:
    branch:
    - test
```