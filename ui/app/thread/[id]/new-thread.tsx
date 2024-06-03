import PromptInput from "@/app/thread/[id]/prompt-input";

export default function NewThread() {
  return (
    <>
      <span className="w-full text-center mb-2">
        <h1>What do you want to explore?</h1>
      </span>
      <PromptInput threadId="" />
    </>
  );
}
