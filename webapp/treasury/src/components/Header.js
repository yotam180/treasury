import { AppBar, makeStyles, Toolbar, Typography } from "@material-ui/core";
import React from "react";
import { Link } from "react-router-dom";

import logo from "../assets/logo.png";

const useStyles = makeStyles((theme) => ({
  appBar: {
    background: theme.palette.background.paper,
    color: theme.palette.text.primary,
  },
  logo: {
    height: "3rem",
    marginRight: theme.spacing(2),
  },
}));

export function Header() {
  const styles = useStyles();

  return (
    <div>
      <AppBar position="static" variant="elevation" className={styles.appBar}>
        <Toolbar>
          <Link to="/">
            <img src={logo} className={styles.logo} alt="Treasury logo" />
          </Link>
          <Typography variant="h5">Treasury</Typography>
        </Toolbar>
      </AppBar>
    </div>
  );
}
