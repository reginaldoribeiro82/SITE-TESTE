Diagrama IHC (Protótipo) - Domínio Notícias

Este documento identifica os principais eventos de interação (ações do usuário) nas telas do domínio "Notícias" do aplicativo "Minha UFG" para o seguinte caso de uso:

Consultar Notícias (UC_Noticias_2)

Tela 1: Menu Principal "Minha UFG"

![Tela Principal](https://github.com/user-attachments/assets/971702b6-f7fb-42bb-ab15-5c192031f043)


Esta é a tela principal do aplicativo, onde o usuário seleciona um dos 8 módulos. O fluxo se inicia aqui.

Eventos de Interação Identificados:
______________________________________________________________________________________________________________________________________________________________

Evento Nº 01: Acessar o Domínio Notícias

Evento: Ação ao Tocar

Componente: Ícone/Botão "Notícias" (no menu principal de 8 ícones).

Resultado Esperado: O aplicativo fecha a Tela Principal e abre a "Tela 2: Notícias", carregando as notícias da regional padrão do usuário.

______________________________________________________________________________________________________________________________________________________________

Tela 2: Tela "Notícias"

![ITela Noticias](https://github.com/user-attachments/assets/7991c46b-883b-4f59-a1d7-a326d1440feb)


Esta tela é exibida após o Evento Nº 01. É aqui que o usuário pode filtrar notícias por regional e selecionar uma para ler o conteúdo completo.

Eventos de Interação Identificados:
______________________________________________________________________________________________________________________________________________________________

Evento Nº 02: Selecionar a Regional

Evento: Ação ao Tocar

Componente: Botão/Seletor de Regional (localizado no topo da tela, ao lado do título "Notícias").

Resultado Esperado: Exibe um menu suspenso (dropdown) permitindo ao usuário escolher uma regional diferente. Ao selecionar uma nova opção, a lista de notícias abaixo é 
recarregada.
______________________________________________________________________________________________________________________________________________________________

Evento Nº 03: Clicar no resumo da notícia

Evento: Ação ao Tocar

Componente: O "card" completo da notícia (a área retangular que contém a imagem, o título e o texto resumo).

Resultado Esperado: O aplicativo solicita ao Sistema Operacional a abertura do navegador padrão para exibir a página web da UFG com a notícia completa.
______________________________________________________________________________________________________________________________________________________________

Evento Nº 04: Retornar ao Menu Principal

Evento: Ação ao Tocar

Componente: Ícone "Voltar" (seta para a esquerda, no canto superior esquerdo).

Resultado Esperado: O aplicativo fecha a "Tela 2: Notícias" e exibe novamente a "Tela 1: Menu Principal".
