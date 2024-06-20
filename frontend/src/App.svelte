<script lang="ts">
  import { onMount } from "svelte";
  import Sidebar from "./Sidebar.svelte";
  import Settings from "./settings/page.svelte";
  import Page from "./home/page.svelte";
  import { BrowserOpenURL } from "../wailsjs/runtime";
  import { IsSignedIn } from "../wailsjs/go/main/App";
  import { toDialog } from "./shared";
  let online;
  let all;
  let pending;
  function make1Active() {
    online.classList.add("tab-active");
    all.classList.remove("tab-active");
    pending.classList.remove("tab-active");
  }
  function make2Active() {
    online.classList.remove("tab-active");
    all.classList.add("tab-active");
    pending.classList.remove("tab-active");
  }
  function make3Active() {
    online.classList.remove("tab-active");
    all.classList.remove("tab-active");
    pending.classList.add("tab-active");
  }
  import SignIn from "./SignIn.svelte";
  import { inSettingsPage, isCardOpen } from "./store";
  import UserClick from "./UserClick.svelte";
  let loggedIn = false;
  onMount(async () => {
    loggedIn = await IsSignedIn();
    document.querySelectorAll("a").forEach((link) => {
      const url = link.href;
      link.addEventListener("click", (e) => {
        if (!url.startsWith("http://#") && !url.startsWith("file://")) {
          e.preventDefault();
          BrowserOpenURL(link.href);
        }
      });
    });
    document.addEventListener("click", (e) => {
      let el = e.target as HTMLElement;
      let close = true;
      while (true) {
        if (el.parentElement.tagName == "BODY") {
          break;
        } else if (
          el.className.includes("part_of_user") ||
          el.parentElement.className.includes("part_of_user")
        ) {
          close = false;
          break;
        }
        el = el.parentElement;
      }
      if (close) {
        isCardOpen.set(false);
      }
    });
  });
</script>

{#if $inSettingsPage && loggedIn}
  <Settings />
{:else if loggedIn}
  <div class="flex flex-row h-full font-mono">
    <Sidebar></Sidebar>
    <Page></Page>
    <div class="flex-grow p-4 bg-base-200">
      <div class="flex flex-row gap-4 items-center">
        <div
          class="flex gap-2 items-center h-12 text-sm font-bold rounded-lg text-greyish w-fit neutral-content"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="currentColor"
            class="size-6"
          >
            <path
              fill-rule="evenodd"
              d="M8.25 6.75a3.75 3.75 0 1 1 7.5 0 3.75 3.75 0 0 1-7.5 0ZM15.75 9.75a3 3 0 1 1 6 0 3 3 0 0 1-6 0ZM2.25 9.75a3 3 0 1 1 6 0 3 3 0 0 1-6 0ZM6.31 15.117A6.745 6.745 0 0 1 12 12a6.745 6.745 0 0 1 6.709 7.498.75.75 0 0 1-.372.568A12.696 12.696 0 0 1 12 21.75c-2.305 0-4.47-.612-6.337-1.684a.75.75 0 0 1-.372-.568 6.787 6.787 0 0 1 1.019-4.38Z"
              clip-rule="evenodd"
            />
            <path
              d="M5.082 14.254a8.287 8.287 0 0 0-1.308 5.135 9.687 9.687 0 0 1-1.764-.44l-.115-.04a.563.563 0 0 1-.373-.487l-.01-.121a3.75 3.75 0 0 1 3.57-4.047ZM20.226 19.389a8.287 8.287 0 0 0-1.308-5.135 3.75 3.75 0 0 1 3.57 4.047l-.01.121a.563.563 0 0 1-.373.486l-.115.04c-.567.2-1.156.349-1.764.441Z"
            />
          </svg>
          Friends
        </div>
        <div role="tablist" class="tabs tabs-boxed">
          <button
            role="tab"
            class="tab tab-active"
            bind:this={online}
            on:click={make1Active}>Online</button
          >
          <button role="tab" class="tab" bind:this={all} on:click={make2Active}
            >All</button
          >
          <button
            role="tab"
            class="tab"
            bind:this={pending}
            on:click={make3Active}>Pending</button
          >
        </div>
        <button class="btn btn-secondary btn-sm"
          ><svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="currentColor"
            class="size-4"
          >
            <path
              d="M5.25 6.375a4.125 4.125 0 1 1 8.25 0 4.125 4.125 0 0 1-8.25 0ZM2.25 19.125a7.125 7.125 0 0 1 14.25 0v.003l-.001.119a.75.75 0 0 1-.363.63 13.067 13.067 0 0 1-6.761 1.873c-2.472 0-4.786-.684-6.76-1.873a.75.75 0 0 1-.364-.63l-.001-.122ZM18.75 7.5a.75.75 0 0 0-1.5 0v2.25H15a.75.75 0 0 0 0 1.5h2.25v2.25a.75.75 0 0 0 1.5 0v-2.25H21a.75.75 0 0 0 0-1.5h-2.25V7.5Z"
            />
          </svg>
          Add Friend</button
        >
      </div>
    </div>
  </div>

  <dialog id="status_dialog" class="modal">
    <form
      on:submit|preventDefault={() =>
        toDialog(document.getElementById("status_dialog")).close()}
      class="modal-box"
    >
      <h3 class="mb-2 text-lg font-bold text-center">Set A Status</h3>
      <input
        type="text"
        class="w-full input input-bordered"
        placeholder="Whats up?"
      />
      <div class="modal-action">
        <form method="dialog">
          <button class="btn">Close</button>
        </form>
        <button class="btn btn-primary">Save</button>
      </div>
    </form>
  </dialog>
  <UserClick />
{:else}
  <SignIn />
{/if}
