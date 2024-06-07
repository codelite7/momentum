import PromptInput from "@/app/thread/[id]/prompt-input";
import ParentThread from "@/app/thread/[id]/parent-thread";
import ThreadHeader from "@/components/thread/thread-header";

interface NewThreadProps {
  parentId?: string;
}

export default function NewThread({ parentId }: NewThreadProps) {
  return (
    <>
      <ThreadHeader />
      <div className="w-full h-full flex gap-2 px-2 overflow-hidden pb-2">
        {/* header */}
        {parentId && <ParentThread parentId={parentId} />}

        {/* prompt input */}
        <div className="flex flex-col w-full h-full place-content-center p-2">
          <span className="w-full text-center mb-2">
            <h1>What do you want to explore?</h1>
          </span>
          <PromptInput threadId="" />
        </div>
      </div>
    </>
  );
}
