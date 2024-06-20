<script>
  import { onMount } from "svelte";
  import { themeChange } from "theme-change";

  onMount(() => {
    themeChange(false);
  });
</script>

<h1 class="mb-4 text-xl font-bold leading-none text-base-content">Themes</h1>
<div class="grid grid-cols-2 gap-4 rounded-box sm:grid-cols-3 md:grid-cols-4">
  <div
    class="overflow-hidden rounded-lg border border-base-content/20 hover:border-base-content/40 outline outline-2 outline-offset-2 outline-transparent"
    data-set-theme="light"
    data-act-class="!outline-base-content"
  >
    <div
      data-theme="light"
      class="w-full font-sans cursor-pointer bg-base-100 text-base-content"
    >
      <div class="grid grid-cols-5 grid-rows-3">
        <div class="col-start-1 row-span-2 row-start-1 bg-base-200"></div>
        <div class="col-start-1 row-start-3 bg-base-300"></div>
        <div
          class="flex flex-col col-span-4 col-start-2 row-span-3 row-start-1 gap-1 p-2 bg-base-100"
        >
          <div class="font-bold">light</div>
          <div class="flex flex-wrap gap-1">
            <div
              class="flex justify-center items-center w-5 rounded bg-primary aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-primary-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-secondary aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-secondary-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-accent aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-accent-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-neutral aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-neutral-content">A</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div
    class="overflow-hidden rounded-lg border border-base-content/20 hover:border-base-content/40 outline outline-2 outline-offset-2 outline-transparent"
    data-set-theme="pastel"
    data-act-class="!outline-base-content"
  >
    <div
      data-theme="pastel"
      class="w-full font-sans cursor-pointer bg-base-100 text-base-content"
    >
      <div class="grid grid-cols-5 grid-rows-3">
        <div class="col-start-1 row-span-2 row-start-1 bg-base-200"></div>
        <div class="col-start-1 row-start-3 bg-base-300"></div>
        <div
          class="flex flex-col col-span-4 col-start-2 row-span-3 row-start-1 gap-1 p-2 bg-base-100"
        >
          <div class="font-bold">pastel</div>
          <div class="flex flex-wrap gap-1">
            <div
              class="flex justify-center items-center w-5 rounded bg-primary aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-primary-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-secondary aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-secondary-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-accent aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-accent-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-neutral aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-neutral-content">A</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div
    class="overflow-hidden rounded-lg border border-base-content/20 hover:border-base-content/40 outline outline-2 outline-offset-2 outline-transparent"
    data-set-theme="bumblebee"
    data-act-class="!outline-base-content"
  >
    <div
      data-theme="bumblebee"
      class="w-full font-sans cursor-pointer bg-base-100 text-base-content"
    >
      <div class="grid grid-cols-5 grid-rows-3">
        <div class="col-start-1 row-span-2 row-start-1 bg-base-200"></div>
        <div class="col-start-1 row-start-3 bg-base-300"></div>
        <div
          class="flex flex-col col-span-4 col-start-2 row-span-3 row-start-1 gap-1 p-2 bg-base-100"
        >
          <div class="font-bold">bumblebee</div>
          <div class="flex flex-wrap gap-1">
            <div
              class="flex justify-center items-center w-5 rounded bg-primary aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-primary-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-secondary aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-secondary-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-accent aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-accent-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-neutral aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-neutral-content">A</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div
    class="overflow-hidden rounded-lg border border-base-content/20 hover:border-base-content/40 outline outline-2 outline-offset-2 outline-transparent"
    data-set-theme="emerald"
    data-act-class="!outline-base-content"
  >
    <div
      data-theme="emerald"
      class="w-full font-sans cursor-pointer bg-base-100 text-base-content"
    >
      <div class="grid grid-cols-5 grid-rows-3">
        <div class="col-start-1 row-span-2 row-start-1 bg-base-200"></div>
        <div class="col-start-1 row-start-3 bg-base-300"></div>
        <div
          class="flex flex-col col-span-4 col-start-2 row-span-3 row-start-1 gap-1 p-2 bg-base-100"
        >
          <div class="font-bold">emerald</div>
          <div class="flex flex-wrap gap-1">
            <div
              class="flex justify-center items-center w-5 rounded bg-primary aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-primary-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-secondary aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-secondary-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-accent aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-accent-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-neutral aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-neutral-content">A</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div
    class="overflow-hidden rounded-lg border border-base-content/20 hover:border-base-content/40 outline outline-2 outline-offset-2 outline-transparent"
    data-set-theme="cyberpunk"
    data-act-class="!outline-base-content"
  >
    <div
      data-theme="cyberpunk"
      class="w-full font-sans cursor-pointer bg-base-100 text-base-content"
    >
      <div class="grid grid-cols-5 grid-rows-3">
        <div class="col-start-1 row-span-2 row-start-1 bg-base-200"></div>
        <div class="col-start-1 row-start-3 bg-base-300"></div>
        <div
          class="flex flex-col col-span-4 col-start-2 row-span-3 row-start-1 gap-1 p-2 bg-base-100"
        >
          <div class="font-bold">cyberpunk</div>
          <div class="flex flex-wrap gap-1">
            <div
              class="flex justify-center items-center w-5 rounded bg-primary aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-primary-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-secondary aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-secondary-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-accent aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-accent-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-neutral aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-neutral-content">A</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div
    class="border-base-content/20 hover:border-base-content/40 overflow-hidden rounded-lg border outline outline-2 outline-offset-2 outline-transparent !outline-base-content"
    data-set-theme="dark"
    data-act-class="!outline-base-content"
  >
    <div
      data-theme="dark"
      class="w-full font-sans cursor-pointer bg-base-100 text-base-content"
    >
      <div class="grid grid-cols-5 grid-rows-3">
        <div class="col-start-1 row-span-2 row-start-1 bg-base-200"></div>
        <div class="col-start-1 row-start-3 bg-base-300"></div>
        <div
          class="flex flex-col col-span-4 col-start-2 row-span-3 row-start-1 gap-1 p-2 bg-base-100"
        >
          <div class="font-bold">dark</div>
          <div class="flex flex-wrap gap-1">
            <div
              class="flex justify-center items-center w-5 rounded bg-primary aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-primary-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-secondary aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-secondary-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-accent aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-accent-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-neutral aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-neutral-content">A</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div
    class="overflow-hidden rounded-lg border border-base-content/20 hover:border-base-content/40 outline outline-2 outline-offset-2 outline-transparent"
    data-set-theme="forest"
    data-act-class="!outline-base-content"
  >
    <div
      data-theme="forest"
      class="w-full font-sans cursor-pointer bg-base-100 text-base-content"
    >
      <div class="grid grid-cols-5 grid-rows-3">
        <div class="col-start-1 row-span-2 row-start-1 bg-base-200"></div>
        <div class="col-start-1 row-start-3 bg-base-300"></div>
        <div
          class="flex flex-col col-span-4 col-start-2 row-span-3 row-start-1 gap-1 p-2 bg-base-100"
        >
          <div class="font-bold">forest</div>
          <div class="flex flex-wrap gap-1">
            <div
              class="flex justify-center items-center w-5 rounded bg-primary aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-primary-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-secondary aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-secondary-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-accent aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-accent-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-neutral aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-neutral-content">A</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div
    class="overflow-hidden rounded-lg border border-base-content/20 hover:border-base-content/40 outline outline-2 outline-offset-2 outline-transparent"
    data-set-theme="dracula"
    data-act-class="!outline-base-content"
  >
    <div
      data-theme="dracula"
      class="w-full font-sans cursor-pointer bg-base-100 text-base-content"
    >
      <div class="grid grid-cols-5 grid-rows-3">
        <div class="col-start-1 row-span-2 row-start-1 bg-base-200"></div>
        <div class="col-start-1 row-start-3 bg-base-300"></div>
        <div
          class="flex flex-col col-span-4 col-start-2 row-span-3 row-start-1 gap-1 p-2 bg-base-100"
        >
          <div class="font-bold">dracula</div>
          <div class="flex flex-wrap gap-1">
            <div
              class="flex justify-center items-center w-5 rounded bg-primary aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-primary-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-secondary aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-secondary-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-accent aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-accent-content">A</div>
            </div>
            <div
              class="flex justify-center items-center w-5 rounded bg-neutral aspect-square lg:w-6"
            >
              <div class="text-sm font-bold text-neutral-content">A</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
