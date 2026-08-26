## Why

Alternar manualmente configurações globais de diferentes AI coding harnesses torna difícil manter perfis simultâneos, isolados e compreensíveis. Precisamos de um control plane local que resolva capabilities, materialize configurações nativas e supervisione sessões persistentes sem substituir o comportamento interno de cada harness.

## What Changes

- Introduz um registry declarativo de profiles, capabilities, resources e security contexts, com fontes locais e Git.
- Resolve profiles de forma determinística, preservando provenance, detectando conflitos e separando funcionalidade de autoridade.
- Introduz adapters externos e versionados para detectar harnesses, validar compatibilidade e produzir planos declarativos de materialização.
- Materializa uma runtime instance independente por sessão, com configuração específica do vendor, environment filtrado e secrets resolvidos somente no spawn.
- Introduz um daemon local que mantém processos, PTYs e sessões ativos quando a CLI/TUI é fechada.
- Oferece CLI e TUI para criar, listar, inspecionar, anexar, desanexar e encerrar sessões simultâneas.
- Respeita configurações nativas presentes no projeto e as inclui no diagnóstico de configuração efetiva.
- Solicita consentimento antes de operar em um repositório ainda não confiado.
- Detecta alterações feitas pelo harness em arquivos gerenciados e as oferece como diff, sem promoção automática.
- Fornece adapters oficiais para Claude Code, OpenAI Codex e GitHub Copilot CLI pelo mesmo protocolo disponível a adapters de terceiros.
- Não inclui no MVP recuperação de sessões após falha do daemon ou reboot, sandbox forte nem marketplace de adapters; esses itens ficam registrados como evolução futura.

## Capabilities

### New Capabilities

- `profile-resolution`: Definição e resolução de profiles, capabilities, resources e security contexts provenientes de registries locais e Git.
- `tool-adapters`: Descoberta, confiança e execução de adapters externos versionados que traduzem runtimes para configurações nativas dos harnesses.
- `runtime-materialization`: Criação segura e inspecionável de runtime instances isoladas por sessão, incluindo environment, secrets references e detecção de mudanças.
- `persistent-sessions`: Supervisão local persistente de processos e PTYs, permitindo múltiplas sessões que sobrevivem ao fechamento da interface.
- `terminal-workspace`: Operação das sessões por CLI e TUI, incluindo criação, inspeção, attach, detach e encerramento.
- `repository-trust`: Consentimento e diagnóstico para operação em repositórios ainda não confiados, preservando configurações nativas do projeto.

### Modified Capabilities

Nenhuma.

## Impact

- Novo binário Go com modos cliente e daemon local.
- Novo protocolo IPC local entre CLI/TUI e daemon.
- Novo protocolo de subprocesso para adapters oficiais e de terceiros.
- Persistência local de registry, runtime instances, metadados de sessão e decisões de confiança.
- Integração inicial com Claude Code, OpenAI Codex e GitHub Copilot CLI em Linux e macOS.
- Dependência de suporte a PTY por plataforma; sem dependência obrigatória de tmux.
