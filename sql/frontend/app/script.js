import { reactive, html } from 'https://esm.sh/@arrow-js/core';

const state = reactive({
	creators: [],
})

function creatorTemplate(creator) {
	return html`
		<tr>
			<td>${creator.ID}</td>
			<td>${creator.Username}</td>
			<td>${creator.FirstName}</td>
			<td>${creator.LastName}</td>
		</tr>
	`.key(creator.id)
}

const appTemplate = html`
<h1>Search Creators</h1>
<p>Powered by Go and ArrowJS</p>
<input id="search-field" type="text" name="search" />
<button @click="${getCreator}">Find</button>
<table>
	<tr>
		<th>ID</th>
		<th>Username</th>
		<th>FirstName</th>
		<th>LastName</th>
	</tr>
  ${() => state.creators.map(creatorTemplate)}
</table>
`;
const appElement = document.getElementById('app');
appTemplate(appElement);

async function getCreator() {
	const search = document.querySelector("#search-field").value;
	const response = await fetch("/api/creators/" + search);
  const creators = await response.json();
	state.creators = creators;
}
