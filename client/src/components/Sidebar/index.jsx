import React from "react";
import SidebarItem from "../SidebarItem";
import * as styles from "./index.module.css";
import propTypes from "prop-types";
import { tabs } from "../../constants/sidebar";
// import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
// import { faUserLock } from "@fortawesome/free-solid-svg-icons";

const Sidebar = ({ tab, setTab }) => {
  const onClick = (id) => {
    setTab(id);
  };

  return (
    <div className={styles.sidebar}>
      {/* <div className={styles.logo}>
        <FontAwesomeIcon icon={faUserLock} />
      </div> */}
      <div className={styles.sideBarItems}>
        {tabs.map(({ name, id, icon }) => (
          <SidebarItem
            key={id}
            text={name}
            selected={tab === id}
            icon={icon}
            id={id}
            onClick={onClick}
          />
        ))}
      </div>
    </div>
  );
};

Sidebar.propTypes = {
  tab: propTypes.string.isRequired,
  setTab: propTypes.func.isRequired,
};

export default Sidebar;
