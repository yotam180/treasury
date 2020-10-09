import { AppBar, makeStyles, Toolbar, Typography } from "@material-ui/core";
import React from "react";

const useStyles = makeStyles((theme) => ({
  appBar: {
    background: theme.palette.background.paper,
    color: theme.palette.text.primary,
  },
}));

export function Header() {
  const styles = useStyles();

  return (
    <div>
      <AppBar position="static" variant="dense" className={styles.appBar}>
        <Toolbar>
          <Typography variant="h5">Treasury Artifact Repository</Typography>
        </Toolbar>
      </AppBar>
    </div>
  );
}
