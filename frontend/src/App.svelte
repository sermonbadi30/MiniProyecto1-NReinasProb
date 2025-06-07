<script>
  let nReinas = 0;
  let tamano = 0;
  let consola = [];

  function generarTablero() {
    consola = [...consola, `Generando tablero de ${tamano}x${tamano} con ${nReinas} reinas...`];
  }
</script>

<style>
  :global(body) {
    margin: 0;
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
    background-color: #121212;
    color: #e0e0e0;
  }

  .container {
    display: flex;
    height: 100vh;
  }

  .sidebar {
    width: 320px;
    background: #1e1e1e;
    padding: 20px;
    box-sizing: border-box;
    display: flex;
    flex-direction: column;
    gap: 20px;
    border-right: 1px solid #333;
  }

  .content {
    flex: 1;
    padding: 30px;
    background: #252525;
    box-sizing: border-box;
  }

  .input-group {
    display: flex;
    flex-direction: column;
  }

  label {
    font-weight: 600;
    margin-bottom: 8px;
    font-size: 0.95rem;
  }

  input {
    background: #333;
    border: none;
    border-radius: 5px;
    color: #eee;
    padding: 10px;
    font-size: 1rem;
    outline-offset: 2px;
    outline-color: #007acc;
  }

  input:focus {
    outline-style: solid;
  }

  button {
    background: #007acc;
    border: none;
    padding: 12px;
    border-radius: 6px;
    color: white;
    font-weight: 600;
    cursor: pointer;
    transition: background-color 0.3s ease;
  }

  button:hover {
    background: #005f99;
  }

  .consola-label {
    font-weight: 700;
    margin-bottom: 6px;
    font-size: 1rem;
  }

  .consola {
    background: #0f111a;
    color: #64ffda;
    font-family: 'Source Code Pro', monospace, monospace;
    padding: 15px;
    height: 200px;
    border-radius: 8px;
    overflow-y: auto;
    white-space: pre-wrap;
    box-shadow: inset 0 0 10px #004d40;
    user-select: text;
  }
</style>

<div class="container">
  <aside class="sidebar">
    <div class="input-group">
      <label for="nReinas">Cantidad de reinas</label>
      <input id="nReinas" type="number" min="1" bind:value={nReinas} />
    </div>

    <div class="input-group">
      <label for="tamano">Tamaño del tablero (NxN)</label>
      <input id="tamano" type="number" min="1" bind:value={tamano} />
    </div>

    <button on:click={generarTablero}>Generar tablero</button>

    <div class="input-group" style="margin-top: auto;">
      <p class="consola-label">Consola</p>
      <div class="consola">
        {#each consola as linea}
          {linea}
          <br />
        {/each}
      </div>
    </div>
  </aside>

  <main class="content">
    <h2>Espacio para el tablero</h2>
    <!-- Aquí se renderizará el tablero -->
  </main>
</div>
