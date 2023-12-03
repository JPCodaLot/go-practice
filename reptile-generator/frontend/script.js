import { reactive, html } from 'https://esm.sh/@arrow-js/core';

const state = reactive({
	reptile: null,
})

function reptileTemplate(reptile) {
	if (reptile === null) {
		return
	}
	return html`
		<article class="card">
			<h2 class="title">${reptile.name}</h2>
			<p class="subtitle">${reptile.latin_name}</p>
			<img class="photo" src="${reptile.photo}" />
			<iframe width="560" height="315" src="https://www.youtube-nocookie.com/embed/${reptile.video}" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>
		</aticle>`
}

const appTemplate = html`
<h1 class="title">Reptile Generator</h1>
<p class="subtitle">Powered by Go, Redis, and ArrowJS</p>
<header class="button-bar">
	<button class="button" @click="${generateReptile}">Generate</button>
</header>
${() => reptileTemplate(state.reptile)}
`;
const appElement = document.getElementById('app');
appTemplate(appElement);

async function generateReptile() {
	const response = await fetch("/api/reptiles/pick");
  state.reptile = await response.json();
}
