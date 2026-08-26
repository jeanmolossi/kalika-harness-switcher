## Why

O projeto ainda não possui uma base Go compilável e verificável sobre a qual as capabilities do control plane possam evoluir incrementalmente. Um walking skeleton mínimo estabelece identidade, toolchain, testes e CI sem antecipar abstrações do domínio ou adicionar dependências prematuras.

## What Changes

- Inicializa o módulo Go `github.com/jeanmolossi/kalika-harness-switcher`.
- Cria um único binário público chamado `khs`.
- Introduz uma CLI mínima baseada exclusivamente na biblioteca padrão, com ajuda e informação de versão.
- Mantém `main` sem lógica de aplicação e torna a execução da CLI testável através de argumentos, writers e exit code.
- Adiciona metadata de build para versão, commit e data, com defaults determinísticos para desenvolvimento e suporte a injeção por linker flags.
- Fixa Go 1.26.0 como requisito mínimo e Go 1.26.7 como toolchain preferida.
- Adiciona comandos de desenvolvimento para format check, vet, testes, race detector e build.
- Adiciona CI para build e testes em Linux e macOS.
- Estabelece que paths devem ser construídos com `filepath` e partir de resolvedores semânticos nativos do Go, nunca de convenções inferidas ou caminhos hardcoded.
- Estabelece que comportamentos específicos de sistema operacional devem ficar isolados em implementações explícitas por plataforma, com Linux e macOS cobertos inicialmente e Windows preservado como evolução futura.
- Documenta o propósito, estado experimental, requisitos e fluxo básico de desenvolvimento.
- Não cria antecipadamente packages de domínio, daemon, IPC, storage, adapters, PTY ou TUI.
- Não adiciona dependências Go externas.

## Capabilities

### New Capabilities

Nenhuma. Esta change estabelece somente tooling e estrutura de desenvolvimento, portanto opta por não criar specs de comportamento.

### Modified Capabilities

Nenhuma.

## Impact

- Novos arquivos de módulo, entry point, build information, CLI mínima, testes, Makefile e workflow de CI.
- O comando público e futuros exemplos passam a usar `khs`.
- Desenvolvedores precisam de uma instalação compatível com a seleção automática de toolchain do Go ou da toolchain preferida instalada localmente.
- Todo código futuro que manipule paths ou varie por OS fica sujeito à convenção multiplataforma estabelecida nesta base.
- Não há migração nem breaking change porque ainda não existe aplicação distribuída.
