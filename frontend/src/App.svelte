<script>
  import { SetFolderPath } from "../wailsjs/go/main/App.js";
  import Gallery from "./Gallery.svelte";
  import { imageFolderPath, updateFolderPath } from "./store/imageFolder.js";

  /* 
  #E8C872 - yellow
  #FFF3CF - lite yellow
  #C9D7DD - lite blue
  #637A9F - blue
  */

  let resultText = "Folder with reactions here 👇";
  let imageFolder;

  imageFolderPath.subscribe((value) => {
    imageFolder = value;
  });

  function setFolderPath() {
    SetFolderPath().then((res) => {
      updateFolderPath(res);
    });
  }
</script>

<main>
  <div>Click on tags to delete them</div>
  <div class="result" id="result">{resultText}</div>
  <div class="input-box" id="input">
    <input
      autocomplete="off"
      disabled
      bind:value={imageFolder}
      class="input"
      id="name"
      type="text"
    />
    <button class="btn" on:click={setFolderPath}>Show Images</button>
  </div>
  <Gallery {imageFolder} />
</main>
