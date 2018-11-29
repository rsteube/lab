if [ "$(lab version | grep lab)" != "$(cat "${0:h}/_lab.version" 2>/dev/null)" ]; then
  lab completion zsh > "${0:h}/_lab"
  lab version | grep lab > "${0:h}/_lab.version"
fi
