Diagrama IHC (Protótipo) - Domínio Eventos

Este documento identifica os principais eventos de interação (ações do usuário) nas telas do domínio "Eventos" do aplicativo "Minha UFG" para o seguinte caso de uso:

Consultar Eventos (UC_Eventos_5)

Tela 1: Menu Principal "Minha UFG"

![Menu Principal](https://github.com/user-attachments/assets/b0a80a68-d0a2-48eb-889f-7aa6707e0265)


Esta é a tela principal do aplicativo, onde o usuário seleciona um dos 8 domínio. O fluxo se inicia aqui.
_________________________________________________________________________________________________________________________________________________________

Eventos de Interação Identificados:

Evento Nº 01: Acessar o Domínio Eventos

Evento: Ação ao Tocar

Componente: Ícone/Botão "Eventos" (no menu principal de 8 ícones).

Resultado Esperado: O aplicativo fecha a Tela Principal e abre a "Tela 2: Eventos", carregando os filtros e eventos da regional e data atuais.
_________________________________________________________________________________________________________________________________________________________

Tela 2: Tela "Eventos"

![Tela de Eventos](https://github.com/user-attachments/assets/27c35892-1dc2-4a3e-85fd-0e9ddc74b74c)


Esta tela é exibida após o Evento Nº 01. É aqui que o usuário pode filtrar eventos por regional, data e selecionar um para ver os detalhes.

_________________________________________________________________________________________________________________________________________________________

Eventos de Interação Identificados:

Evento Nº 02: Selecionar a Regional

Evento: Ação ao Tocar

Componente: Botão/Seletor de Regional (localizado no topo da tela, ex: "Regional Goiânia").

Resultado Esperado: Exibe um menu suspenso (dropdown) permitindo ao usuário escolher uma regional diferente. Ao selecionar uma nova opção, a lista de eventos abaixo é recarregada.

_________________________________________________________________________________________________________________________________________________________

Evento Nº 03: Selecionar uma Data

Evento: Ação ao Tocar

Componente: Um dia no componente de "Calendário".

Resultado Esperado: A lista de eventos abaixo é atualizada para mostrar apenas os eventos que ocorrem na data selecionada.

_________________________________________________________________________________________________________________________________________________________
Evento Nº 04: Clicar em um Evento

Evento: Ação ao Tocar

Componente: O "card" de um evento na lista (a área que contém a miniatura, nome e data).

Resultado Esperado: O aplicativo solicita ao Sistema Operacional a abertura do navegador padrão para exibir a página web da UFG com os detalhes do evento.
_________________________________________________________________________________________________________________________________________________________

Evento Nº 05: Retornar ao Menu Principal

Evento: Ação ao Tocar

Componente: Ícone "Voltar" (seta para a esquerda, no canto superior esquerdo).

Resultado Esperado: O aplicativo fecha a "Tela 2: Eventos" e exibe novamente a "Tela 1: Menu Principal".
