import { createMuiTheme } from "@material-ui/core";

export const theme = createMuiTheme({
  palette: {
    type: "dark",
    primary: { main: "rgba(118, 255, 3, 0.8)" },
    primary2: { main: "#80d8ff" },
    secondary: { main: "rgba(174, 234, 0, 0.83)" },
    accent2: { main: "#1b5e20" },
    text: {
      primary: "#fff",
    },
    accent3: { main: "rgba(129, 199, 132, 0.51)" },
  },
});
