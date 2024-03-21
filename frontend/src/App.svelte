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
  let childRef;
  let newImageLink = "";

  imageFolderPath.subscribe((value) => {
    imageFolder = value;
  });

  function setFolderPath() {
    SetFolderPath().then((res) => {
      updateFolderPath(res);
    });
  }

  function addNewImage() {
    if (newImageLink.length > 0) {
      childRef.addNewImage(newImageLink);
      newImageLink = "";
    }
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
    <button class="btn" on:click={setFolderPath}>Select meme folder</button>
    <input bind:value={newImageLink} />
    <button class="btn" on:click={addNewImage}>Add new link</button>
  </div>
  <Gallery {imageFolder} bind:this={childRef} />
</main>
