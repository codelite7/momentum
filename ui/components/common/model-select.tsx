"use client"

import {Select, SelectItem} from "@nextui-org/react";
import { useState } from "react";

interface modelOption {
    label: string
    value: string
}

const models = [
    {label: 'GPT 4', value: 'gpt-4', icon: 'openai.png'},
    {label: 'Opus', value: 'opus', icon: 'anthropic.png'},
    {label: 'Mistral', value: 'mistral', icon: 'mistral.svg'},
    {label: 'Groq', value: 'groq', icon: 'groq.png'},
]

const logos = {
    'gpt-4': 'openai.png',
    opus: 'anthropic.png',
    mistral: 'mistral.svg',
    groq: 'groq.png'
}


export default function ModelSelect() {
    const [selectedLogo, setSelectedLogo] = useState("openai.png")
    return (
        <div className="flex w-36">
            <Select
                size="sm"
                variant="bordered"
                defaultSelectedKeys={['gpt-4']}
                startContent={<img src={`/llm-logos/${selectedLogo}`} />}
                onSelectionChange={(event: any) => updateSelectedLogo(setSelectedLogo, event)}
            >
                {models.map((model) => (
                    <SelectItem key={model.value} value={model.value} startContent={<img src={`/llm-logos/${model.icon}`} alt={model.label} />}>
                        {model.label}
                    </SelectItem>
                ))}
            </Select>
        </div>
)
}

function updateSelectedLogo(setter: any, {anchorKey}: any) {
    console.log(anchorKey);
    // @ts-ignore
    console.log(`setting logo to: ${logos[anchorKey]}`)
    // @ts-ignore
    setter(logos[anchorKey])
}