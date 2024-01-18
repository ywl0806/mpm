# MPM

## Introduction

이 CLI는 통합 개발 환경(IDE) 또는 텍스트 편집기에서 프로젝트를 실행하는 프로세스를 간소화합니다. 이 CLI를 사용하면 명령 코드 [프로젝트 이름]을 실행하여 지정된 프로젝트를 원활하게 시작할 수 있습니다.

## Why

그냥

## Installation

```bash
go install github.com/ywl0806/mpm
```

## Usage

원한다면 `.bashrc`나`.zshrc`에 `export PATH="$HOME/go/bin:$PATH"`를 추가하세요

당신의 프로젝트 디렉토리에 가서 다음 커맨드를 실행하면 질의가 시작됩니다.

```bash
mpm add
```

프로젝트 이름과 해당되는 디렉토리를 설정하고 각각의 디렉토리에 해당하는 커맨드와 옵션을 입력하세요.

```
Project Name: my-project
Choose derectories front, server
Commands for directory [front]: code
Options for directory [front]:
Commands for directory [server]: code
Options for directory [server]:
```

`-d`나 `--deep`옵션을 이용하면 디렉토리의 계층을 재귀적으로 가져옵니다

```bash
# 현재의 디렉토리 2계층 아래까지 가져옴
mpm add -d 2
```

다음 커맨드를 실행하면 가장 최근에 실행되었거나 추가된 프로젝트가 실행됩니다

```bash
mpm
```

다음 커맨드를 실행하면 프로젝트를 질의를 통해 프로젝트를 선택하여 실행할수 있습니다

```bash
mpm e
# or
mpm excure
# or
mpm run
```

```bash
# 현재 등록되어있는 프로젝트의 리스트
mpm list
# 현재 등록되어있는 프로젝트들의 상세정보
mpm list -d
```

```bash
# 프로젝트를 선택하여 삭제
mpm delete
```
