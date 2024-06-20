<script lang="ts">
  import { PageType, settings_page } from "./page";
  type Page = {
    id: PageType;
    name: string;
    active?: boolean;
  };
  export let name: string;
  export let pages: Page[];
</script>

<div class="flex flex-col gap-1">
  <div class="px-3 font-bold text-base-content">{name}</div>
  {#each pages as page}
    <button
      class="justify-start font-normal btn btn-sm btn-ghost text-base-content"
      class:btn-active={$settings_page == page.id}
      on:click={() => {
        settings_page.set(page.id);
        page.active = true;
      }}
    >
      {page.name}
    </button>
  {/each}
</div>
