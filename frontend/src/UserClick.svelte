<script lang="ts">
  import { isCardOpen } from "./store";
  import { toDialog } from "./shared";
  import { SignOut } from "../wailsjs/go/main/App";
  import { WindowReload } from "../wailsjs/runtime/runtime";
</script>

<div
  class="fixed bottom-20 left-24 w-96 rounded-xl card-compact bg-base-200 part_of_user text-base-content"
  class:hidden={!$isCardOpen}
  class:visible={$isCardOpen}
>
  <div class="relative w-full h-12 bg-violet-400 rounded-t-xl">
    <div class="absolute top-4 left-8 avatar">
      <div class="w-16 rounded-full outline outline-8 outline-base-200">
        <img
          src="https://api.dicebear.com/8.x/bottts-neutral/svg?scale=85&seed=1"
          alt="avatar"
        />
      </div>
    </div>
  </div>

  <div class="mt-6 card-body">
    <div class="flex flex-col p-4 rounded-lg bg-base-100">
      <div class="text-lg">
        <h2 class="font-bold leading-none">DisplayName</h2>
        <span class="leading-none"> username </span>
      </div>
      <div class="my-2 mb-4 w-full h-1 rounded-full div bg-base-300"></div>
      <div class="flex flex-col gap-2 leading-none">
        <div class="font-bold">MEMBER SINCE</div>
        <div>23 Nov 2023</div>
      </div>
      <div class="my-4 w-full h-1 rounded-full div bg-base-300"></div>
      <button
        class="flex gap-2 items-center p-2 mb-2 text-left rounded transition active:bg-primary hover:bg-base-200 hover:font-bold text-base-content"
        on:click={() =>
          toDialog(document.getElementById("status_dialog")).showModal()}
        ><svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 20 20"
          fill="currentColor"
          class="size-5"
        >
          <path
            d="M13.92 3.845a19.362 19.362 0 0 1-6.3 1.98C6.765 5.942 5.89 6 5 6a4 4 0 0 0-.504 7.969 15.97 15.97 0 0 0 1.271 3.34c.397.771 1.342 1 2.05.59l.867-.5c.726-.419.94-1.32.588-2.02-.166-.331-.315-.666-.448-1.004 1.8.357 3.511.963 5.096 1.78A17.964 17.964 0 0 0 15 10c0-2.162-.381-4.235-1.08-6.155ZM15.243 3.097A19.456 19.456 0 0 1 16.5 10c0 2.43-.445 4.758-1.257 6.904l-.03.077a.75.75 0 0 0 1.401.537 20.903 20.903 0 0 0 1.312-5.745 2 2 0 0 0 0-3.546 20.902 20.902 0 0 0-1.312-5.745.75.75 0 0 0-1.4.537l.029.078Z"
          />
        </svg>
        Set Status</button
      >
      <button
        class="flex gap-2 items-center p-2 text-left rounded transition active:bg-primary hover:text-white hover:bg-error text-base-content"
        on:click={() => {
          SignOut();
          WindowReload();
        }}
        ><svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
          class="size-5"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M8.25 9V5.25A2.25 2.25 0 0 1 10.5 3h6a2.25 2.25 0 0 1 2.25 2.25v13.5A2.25 2.25 0 0 1 16.5 21h-6a2.25 2.25 0 0 1-2.25-2.25V15m-3 0-3-3m0 0 3-3m-3 3H15"
          />
        </svg>
        Sign Out</button
      >
    </div>
  </div>
</div>
