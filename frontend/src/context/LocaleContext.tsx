import React, { ReactNode, useState } from "react";

import { getLocale } from "locale";

// Context
export interface LocaleContextType {
	setLocale: (locale: string) => void;
	locale?: string;
}

const initalValue = null;
const LocaleContext = React.createContext<LocaleContextType | null>(
	initalValue,
);

// Provider
interface Props {
	children?: ReactNode;
}
function LocaleProvider({ children }: Props) {
	const [locale, setLocaleValue] = useState(getLocale());

	const setLocale = async (locale: string) => {
		setLocaleValue(locale);
	};

	const value = { locale, setLocale };

	return (
		<LocaleContext.Provider value={value}>{children}</LocaleContext.Provider>
	);
}

function useLocaleState() {
	const context = React.useContext(LocaleContext);
	if (!context) {
		throw new Error(`useLocaleState must be used within a LocaleProvider`);
	}
	return context;
}

export { LocaleProvider, useLocaleState };
export default LocaleContext;
