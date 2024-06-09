"use client";

import { Select, SelectItem } from "@nextui-org/react";
import { useState } from "react";
import { useSetAtom } from "jotai/index";

import { modelSelectAtom } from "@/state/atoms";
import { ThreadProvider } from "@/__generated__/graphql";

interface modelOption {
  label: string;
  value: string;
}

const models = [
  { label: "GPT 4", value: "openai", iconExtension: "png" },
  { label: "Claude", value: "anthropic", iconExtension: "png" },
  { label: "Groq", value: "groq", iconExtension: "png" },
];

export default function ModelSelect() {
  const setAtomValue = useSetAtom(modelSelectAtom);
  const [value, setValue] = useState(new Set(["openai"]));

  return (
    <div className="flex w-36">
      <Select
        aria-label="Select a model"
        classNames={{
          base: "max-w-xs",
          trigger: "h-12",
        }}
        items={models}
        renderValue={(items) => {
          return items.map((item) => (
            <div key={item.key} className="flex items-center gap-2">
              <img
                alt={item.data?.label}
                src={`/llm-logos/${item.data?.value}.${item.data?.iconExtension}`}
              />
              <div className="flex flex-col">
                <span>{item.data?.label}</span>
              </div>
            </div>
          ));
        }}
        selectedKeys={value}
        size="sm"
        onSelectionChange={(e: any) => {
          if (e.size > 0) {
            setValue(e);
            setAtomValue(e.anchorKey as ThreadProvider);
          }
        }}
      >
        {(model) => (
          <SelectItem key={model.value} textValue={model.label}>
            <div className="flex gap-2 items-center">
              <img
                alt={model.label}
                src={`/llm-logos/${model.value}.${model.iconExtension}`}
              />
              <div className="flex flex-col">
                <span className="text-small">{model.label}</span>
              </div>
            </div>
          </SelectItem>
        )}
      </Select>
    </div>
  );
}
