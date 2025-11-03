Diagrama IHC (Protótipo) - Domínio Cardápio RU

Este documento identifica os principais eventos de interação (ações do usuário) nas telas do domínio "Restaurante Universitário" para o seguinte caso de uso:

Consultar Cardápio do RU (UC_RU_4)


Tela 1: Menu Principal "Minha UFG"

Esta é a tela principal do aplicativo, onde o usuário seleciona um dos 8 domínios. O fluxo se inicia aqui.

![Menu Principal](https://github.com/user-attachments/assets/d96aacf7-a3b1-4952-b9d5-0180ff1da863)


Eventos de Interação Identificados:

______________________________________________________________________________________________________________________________________________________________

Evento Nº 01: Acessar o Domínio RU

Evento: Ação ao Tocar

Componente: Ícone/Botão "RU" (no menu principal de 8 ícones).

Resultado Esperado: O aplicativo fecha a Tela Principal e abre a "Tela 2: Menu RU", exibindo as opções de refeição.




__________________________________________________________________________________________________________________________________________________________________
Tela 2: Menu "RU" (Seleção de Refeição)

![Menu RU](https://github.com/user-attachments/assets/dac1f0be-4aa9-4635-96a3-53b6bbe6ce57)

Esta tela é exibida após o Evento Nº 01. Ela mostra os tipos de refeição disponíveis (Café da Manhã, Almoço, Jantar).



Eventos de Interação Identificados:
______________________________________________________________________________________________________________________________________________________________________

Evento Nº 02: Selecionar Tipo de Refeição

Evento: Ação ao Tocar

Componente: Botão "Café da Manhã", "Almoço" ou "Jantar".

Resultado Esperado: O aplicativo abre a "Tela 3: Cardápio", carregando as informações para o tipo de refeição selecionado (ex: "Almoço").
______________________________________________________________________________________________________________________________________________________________________

Evento Nº 03: Retornar ao Menu Principal (Tela 2)

Evento: Ação ao Tocar

Componente: Ícone "Voltar" (seta para a esquerda, no canto superior esquerdo).

Resultado Esperado: O aplicativo fecha a "Tela 2: Menu RU" e exibe novamente a Tela 1 (Menu Principal).
______________________________________________________________________________________________________________________________________________________________________

Tela 3: Tela "Cardápio" (Seleção de Dia)

![Imagem do WhatsApp de 2025-10-31 à(s) 07 36 31_3572db8b](https://github.com/user-attachments/assets/9e0a9162-ed86-4f34-826e-9e6e1d562bda)


Esta tela é exibida após o Evento Nº 02. Ela mostra os dias da semana e o cardápio detalhado da refeição escolhida.

Eventos de Interação Identificados:

______________________________________________________________________________________________________________________________________________________________________

Evento Nº 04: Selecionar o Dia da Semana

Evento: Ação ao Tocar 

Componente: Abas/Botões dos dias da semana (ex: "Seg", "Ter", "Qua").

Resultado Esperado: O aplicativo atualiza a área de conteúdo abaixo para exibir o cardápio correspondente ao dia selecionado.

______________________________________________________________________________________________________________________________________________________________________

Evento Nº 05: Retornar ao Menu "RU" (Tela 3)

Evento: Ação ao Tocar

Componente: Ícone "Voltar" (seta para a esquerda, no canto superior esquerdo).

Resultado Esperado: O aplicativo fecha a "Tela 3: Cardápio" e exibe novamente a "Tela 2: Menu RU".
