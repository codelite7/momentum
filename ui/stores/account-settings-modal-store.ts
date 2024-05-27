import { create } from 'zustand'

type UseAccountSettingsModalStore = {
    isOpen: boolean;
    toggle: () => void;
};

// note the "<MyStore>" next to create
const useAccountSettingsModalStore = create<UseAccountSettingsModalStore>((set) => ({
    isOpen: false,
    toggle: () => set((state) => ({ isOpen: !state.isOpen})),
}))

export default useAccountSettingsModalStore