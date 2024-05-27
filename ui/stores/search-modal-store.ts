import { create } from 'zustand'

type UseSearchModalStore = {
    isOpen: boolean;
    toggle: () => void;
};

// note the "<MyStore>" next to create
const useSearchModalStore = create<UseSearchModalStore>((set) => ({
    isOpen: false,
    toggle: () => set((state) => ({ isOpen: !state.isOpen})),
}))

export default useSearchModalStore