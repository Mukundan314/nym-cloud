import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import * as styles from "./index.module.css";
import propTypes from "prop-types";
import classnames from "classnames";

const SidebarItem = ({ text, selected, icon, id, onClick }) => (
  <div
    role="button"
    className={classnames(styles.sidebarItem, {
      [styles.sidebarItemSelected]: selected,
    })}
    onClick={() => onClick(id)}
  >
    <FontAwesomeIcon icon={icon} />
    <div className={styles.sideBarItemText}>{text}</div>
  </div>
);

SidebarItem.propTypes = {
  text: propTypes.string.isRequired,
  selected: propTypes.bool,
  icon: propTypes.any.isRequired,
  id: propTypes.string.isRequired,
  onClick: propTypes.func.isRequired,
};

SidebarItem.defaultProps = {
  selected: false,
};

export default SidebarItem;
