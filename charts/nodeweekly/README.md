# nodeweekly charts
Helm chart for nodeweeklys

### Installing the charts
From root directory of nodeweekly. Please edit the values.yaml inside charts before applying.
```
helm repo add zufardhiyaulhaq https://charts.zufardhiyaulhaq.com/
helm install zufardhiyaulhaq/nodeweekly --name-template nodeweekly -f values.yaml
```

### Configuration

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| community | string | `"NodeJs Indonesia Community"` |  |
| cronSchedule | string | `"0 8 * * 1"` |  |
| github.branch | string | `"master"` |  |
| github.organization | string | `"zufardhiyaulhaq"` |  |
| github.repository | string | `"community-ops"` |  |
| github.repository_path | string | `"./manifest/nodejs-community/"` |  |
| github.token | string | `"your_token"` |  |
| image.name | string | `"nodeweekly"` |  |
| image.repository | string | `"zufardhiyaulhaq/nodeweekly"` |  |
| image.tag | string | `"0.0.1"` |  |
| image_url | string | `"https://calebmadrigal.com/images/nodejs-logo.png"` |  |
| jobHistoryLimit | int | `1` |  |
| namespace | string | `"nodejs-community"` |  |
| tags | string | `"weekly,nodejs"` |  |

check & modify values.yaml for details
