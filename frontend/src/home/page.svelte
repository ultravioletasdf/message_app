<script>
  import User from "../User.svelte";
</script>

<div class="flex flex-col justify-between items-center p-4 w-64">
  <div
    class="flex flex-col flex-grow gap-2 items-center w-full text-base-content"
  >
    <button class="btn btn-block">
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
      Friends</button
    >
    <div class="text-sm font-bold text-center text-base-content">
      Direct Messages
    </div>
    <button
      class="flex gap-2 items-center p-2 w-full text-base rounded-lg hover:bg-base-200"
    >
      <div class="avatar">
        <div class="w-8 rounded-full">
          <img
            src="https://api.dicebear.com/8.x/bottts-neutral/svg?scale=85&seed=1"
            alt="avatar"
          />
        </div>
      </div>
      Name 1
    </button>
    <button
      class="flex gap-2 items-center p-2 w-full text-base rounded-lg hover:bg-base-200"
    >
      <div class="avatar">
        <div class="w-8 rounded-full">
          <img
            src="https://api.dicebear.com/8.x/bottts-neutral/svg?scale=85&seed=2"
            alt="avatar"
          />
        </div>
      </div>
      Name 2
    </button>
    <button
      class="flex gap-2 items-center p-2 w-full text-base rounded-lg hover:bg-base-200"
    >
      <div class="avatar">
        <div class="w-8 rounded-full">
          <img
            src="https://api.dicebear.com/8.x/bottts-neutral/svg?scale=85&seed=3"
            alt="avatar"
          />
        </div>
      </div>
      Name 3
    </button>
    <button
      class="flex gap-2 items-center p-2 w-full text-base rounded-lg hover:bg-base-200"
    >
      <div class="avatar">
        <div class="w-8 rounded-full">
          <img
            src="https://api.dicebear.com/8.x/bottts-neutral/svg?scale=85&seed=4"
            alt="avatar"
          />
        </div>
      </div>
      Name 4
    </button>
    <button
      class="flex gap-2 items-center p-2 w-full text-base rounded-lg hover:bg-base-200"
    >
      <div class="avatar">
        <div class="w-8 rounded-full">
          <img
            src="https://api.dicebear.com/8.x/bottts-neutral/svg?scale=85&seed=5"
            alt="avatar"
          />
        </div>
      </div>
      Name 5
    </button>
    <button
      class="flex gap-2 items-center p-2 w-full text-base rounded-lg hover:bg-base-200"
    >
      <div class="avatar">
        <div class="w-8 rounded-full">
          <img
            src="https://api.dicebear.com/8.x/bottts-neutral/svg?scale=85&seed=6"
            alt="avatar"
          />
        </div>
      </div>
      Name 6
    </button>
    <button
      class="flex gap-2 items-center p-2 w-full text-base rounded-lg hover:bg-base-200"
    >
      <div class="avatar">
        <div class="w-8 rounded-full">
          <img
            src="https://api.dicebear.com/8.x/bottts-neutral/svg?scale=85&seed=7"
            alt="avatar"
          />
        </div>
      </div>
      Name 7
    </button>
    <button
      class="flex gap-2 items-center p-2 w-full text-base rounded-lg hover:bg-base-200"
    >
      <div class="avatar">
        <div class="w-8 rounded-full">
          <img
            src="https://api.dicebear.com/8.x/bottts-neutral/svg?scale=85&seed=8"
            alt="avatar"
          />
        </div>
      </div>
      Name 8
    </button>
  </div>
  <User></User>
</div>
