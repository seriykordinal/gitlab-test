Запустить gitlab локально на 8080

sudo docker run --detach \           
  --hostname localhost \
  --env GITLAB_OMNIBUS_CONFIG="external_url 'http://localhost:8080'; nginx['listen_port'] = 80; gitlab_rails['gitlab_shell_ssh_port'] = 2222;" \
  --publish 8080:80 --publish 8443:443 --publish 2222:22 \
  --name gitlab \
  --restart always \
  --volume /srv/gitlab/config:/etc/gitlab \
  --volume /srv/gitlab/logs:/var/log/gitlab \
  --volume /srv/gitlab/data:/var/opt/gitlab \
  gitlab/gitlab-ce:latest