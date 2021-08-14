# nodejsweekly

Get data from nodejs news and create Weekly CRDs based on community-operator and push to git datastore

![Version: 1.0.0](https://img.shields.io/badge/Version-1.0.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square) [![made with Go](https://img.shields.io/badge/made%20with-Go-brightgreen)](http://golang.org) [![Github master branch build](https://img.shields.io/github/workflow/status/zufardhiyaulhaq/nodejsweekly/Master)](https://github.com/zufardhiyaulhaq/nodejsweekly/actions/workflows/master.yml) [![GitHub issues](https://img.shields.io/github/issues/zufardhiyaulhaq/nodejsweekly)](https://github.com/zufardhiyaulhaq/nodejsweekly/issues) [![GitHub pull requests](https://img.shields.io/github/issues-pr/zufardhiyaulhaq/nodejsweekly)](https://github.com/zufardhiyaulhaq/nodejsweekly/pulls)

## Installing the Chart

To install the chart with the release name `my-release` and secret `foo`:

```console
kubectl apply -f secret.yaml

helm repo add zufardhiyaulhaq https://charts.zufardhiyaulhaq.com/
helm install nodejsweekly zufardhiyaulhaq/nodejsweekly --values values.yaml --set secret=foo
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| cronSchedule | string | `"0 8 * * 0"` |  |
| image.repository | string | `"zufardhiyaulhaq/nodejsweekly"` |  |
| image.tag | string | `"v1.0.0"` |  |
| secret | string | `""` |  |
| weekly.community | string | `"Node.js Indonesia Community"` |  |
| weekly.image_url | string | `"https://calebmadrigal.com/images/nodejs-logo.png"` |  |
| weekly.namespace | string | `"nodejs-community"` |  |
| weekly.tags | string | `"weekly,nodejs"` |  |

