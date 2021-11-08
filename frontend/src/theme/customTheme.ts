import { theme as chakraTheme, ThemeConfig } from "@chakra-ui/react";
import { extendTheme } from "@chakra-ui/react";

// declare a variable for fonts and set our fonts
const fonts = {
	...chakraTheme.fonts,
	body: `"Source Sans Pro",-apple-system,BlinkMacSystemFont,"Segoe UI",Helvetica,Arial,sans-serif,"Apple Color Emoji","Segoe UI Emoji","Segoe UI Symbol"`,
	heading: `"Source Sans Pro",-apple-system,BlinkMacSystemFont,"Segoe UI",Helvetica,Arial,sans-serif,"Apple Color Emoji","Segoe UI Emoji","Segoe UI Symbol"`,
};

const config: ThemeConfig = {
	initialColorMode: "system",
};

const lightTheme = extendTheme({ fonts, config });
export default lightTheme;
