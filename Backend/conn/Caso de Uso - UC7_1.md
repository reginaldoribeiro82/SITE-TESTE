
Diagrama IHC (Protótipo) - Domínio Locais

Este documento identifica os principais eventos de interação (ações do utilizador) nas telas do domínio "Locais" do aplicativo "Minha UFG" para o seguinte caso de uso:

Consultar Locais no Campus (UC_Locais_7)

Tela 1: Menu Principal "Minha UFG"

![Modulo Menu Principal](https://github.com/user-attachments/assets/04a11d66-aa85-4006-87c0-b8ec40f55949)


Esta é a tela principal do aplicativo, onde o utilizador seleciona um dos 8 domínio. O fluxo inicia-se aqui.
_________________________________________________________________________________________________________________________________________________________

Eventos de Interação Identificados:

Evento Nº 01: Acessar ao Domínio Locais

Evento: Ação ao clicar

Componente: Ícone/Botão "Locais" (no menu principal de 8 ícones).

Resultado Esperado: O aplicativo fecha a Tela Principal e abre a "Tela 2: Locais", carregando a visão "Lista" com os locais do campus/regional padrão.
_________________________________________________________________________________________________________________________________________________________


Tela 2: Tela "Locais"

![Menu Locais](https://github.com/user-attachments/assets/7a73ef0c-bab3-43cf-a469-7f0126f249b5)


Esta tela é exibida após o Evento Nº 01. É aqui que o utilizador pode filtrar locais por campus/regional, alternar entre a visão de lista e mapa, e selecionar um local.

_________________________________________________________________________________________________________________________________________________________


Eventos de Interação Identificados:

Evento Nº 02: Selecionar o Campus/Regional

Evento: Ação ao clicar

Componente: Botão/Seletor de Campus (localizado no topo da tela, ex: "Campus Samambaia").

Resultado Esperado: Exibe um menu suspenso (dropdown) permitindo ao utilizador escolher um campus/regional diferente. A lista (ou mapa) abaixo é recarregada para refletir a nova seleção.
_________________________________________________________________________________________________________________________________________________________

Evento Nº 03: Selecionar um Local (na visão "Lista")

Evento: Ação ao clicar

Componente: Um item na "Lista" (ex: "Biblioteca Central").

Resultado Esperado: O aplicativo requisita os dados desse local e exibe um mapa focado com o marcador (alfinete) nesse local específico.
_________________________________________________________________________________________________________________________________________________________

Evento Nº 04: Retornar ao Menu Principal

Evento: Ação ao clicar

Componente: Ícone "Voltar" (seta para a esquerda, no canto superior esquerdo).

Resultado Esperado: O aplicativo fecha a "Tela 2: Locais" e exibe novamente a "Tela 1: Menu Principal".
_________________________________________________________________________________________________________________________________________________________

Evento Nº 05: Alternar para a Visão "Mapa"

Evento: Ação ao clicar

Componente: Botão/Aba "Mapa" (localizado na parte inferior da tela).

Resultado Esperado: A visão de lista é substituída por um mapa geral (estilo Google Maps) que exibe marcadores (alfinetes) para todos os locais do campus/regional selecionado.

Tela 3: Tela "MAPA"
![Tela MAPA](https://github.com/user-attachments/assets/28e88932-003f-4d82-b534-1db9e048375b)

_________________________________________________________________________________________________________________________________________________________

Evento Nº 06: Iniciar a navegação para o marcador selecionado

Evento: Ação ao clicar

Componente: Marcador colorido circular no mapa

Resultado Esperado: O aplicativo abre um janela para escolher qual aplicativo o usuario gostaria de usar para navegar até o ponto escolhido.
_________________________________________________________________________________________________________________________________________________________

Evento Nº 07: Cancelar a navegação ao marcador selecionado

Evento: Ação ao clicar

Componente: Retangulo " Cancelar"

Resultado Esperado: O aplicativo fecha a janela para escolha da navegação e retorna ao mapa geral que exibe os marcadores.
_________________________________________________________________________________________________________________________________________________________

Evento Nº 08: Retornar ao Menu Principal

Evento: Ação ao clicar

Componente: Ícone "Voltar" (seta para a esquerda, no canto superior esquerdo).

Resultado Esperado: O aplicativo fecha a "Tela 2: Locais" e exibe novamente a "Tela 1: Menu Principal".
