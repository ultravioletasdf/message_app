<script>
  import { onMount } from "svelte";
  import Account from "./account.svelte";
  import Category from "./category.svelte";
  import Theme from "./theme.svelte";
  import { PageType, settings_page } from "./page";
  import { inSettingsPage } from "../store";
  let escPressed = false;

  onMount(() => {
    document.addEventListener("keydown", (e) => {
      if (e.code === "Escape") escPressed = true;
    });
    document.addEventListener("keyup", (e) => {
      if (e.code === "Escape") escPressed = false;
      inSettingsPage.set(false);
    });
  });
</script>

<div class="grid grid-cols-5 h-full">
  <div class="flex col-span-2 justify-end py-16 pr-4 w-full bg-base-100">
    <div class="flex flex-col gap-2 w-48">
      <Category
        name="USER SETTINGS"
        pages={[
          { id: PageType.Account, active: true, name: "My Account" },
          { id: PageType.Connections, name: "Connections" },
        ]}
      />
      <Category
        name="APP SETTINGS"
        pages={[
          { id: PageType.Appearance, name: "Appearance" },
          { id: PageType.Bots, name: "Bots" },
        ]}
      />
    </div>
  </div>
  <div class="col-span-2 px-8 py-16 w-full bg-base-200">
    {#if $settings_page == PageType.Account}
      <Account />
    {:else if $settings_page == PageType.Appearance}
      <Theme></Theme>
    {/if}
  </div>
  <div class="py-16 pl-4 bg-base-200">
    <button class="flex flex-col gap-2 items-center w-fit text-base-content">
      <div
        class="flex justify-center items-center rounded-full transition outline-base-300 text-base-content size-12 bg-base-100 hover:bg-base-300 outline outline-2"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
          class="size-6"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M6 18 18 6M6 6l12 12"
          />
        </svg>
      </div>
      <div class="transition bg-base-100 kbd" class:bg-base-300={escPressed}>
        ESC
      </div>
    </button>
  </div>
</div>
