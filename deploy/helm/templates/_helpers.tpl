{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "krazy-cow.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "krazy-cow.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "krazy-cow.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Common labels
*/}}
{{- define "krazy-cow.labels" -}}
helm.sh/chart: {{ include "krazy-cow.chart" . }}
{{ include "krazy-cow.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{/*
Selector labels
*/}}
{{- define "krazy-cow.selectorLabels" -}}
app.kubernetes.io/name: {{ include "krazy-cow.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{/*
Create the name of the service account to use
*/}}
{{- define "krazy-cow.serviceAccountName" -}}
{{- if .Values.serviceAccount.create -}}
    {{ default (include "krazy-cow.fullname" .) .Values.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.serviceAccount.name }}
{{- end -}}
{{- end -}}

{{/*
Create a comma seperated string with a username and a password
*/}}

{{- define "krazy-cow.getCredentialsString" -}}
{{ if .Values.cowUser -}}
    {{ .Values.cowUser -}}:
{{- else -}}
    {{ randAlphaNum 10 }}:
{{- end -}}
{{- if .Values.cowPassword -}}
    {{- .Values.cowPassword }}
{{- else -}}
    {{- randAlphaNum 10 }}
{{- end -}}
{{- end -}}