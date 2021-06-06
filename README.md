# jasmine

_jasmine is just another service manager_

jasmine is not an init system, nor is it a service supervisor. jasmine is a frontend and control plane for init systems (think `runit`, `openrc` , `s6`, even `systemd`)

jasmine is not a shell command wrapper, and is 100% go. It aims to replace `rsv` and `vsv` as well as `sv` for most things. It is modular (and expandable) by design, but only currently has support for runit planned.

# configuration

## environment variables

| Name                                   | Type/Possible Values | Description                         |
| -------------------------------------- | -------------------- | ----------------------------------- |
| `JASMINE_SUPPRESS_PERMISSIONS_WARNING` | Bool                 | Suppress warnings when UID is not 0 |
