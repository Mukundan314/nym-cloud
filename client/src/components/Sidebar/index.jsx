import React from "react";
import SidebarItem from "../SidebarItem";
import * as styles from "./index.module.css";
import propTypes from "prop-types";
import { tabs } from "../../constants/sidebar";

const Sidebar = ({ tab, setTab }) => (
  <div className={styles.sidebar}>
    {tabs.map(({ name, id, icon }) => (
      <SidebarItem text={name} selected={tab === id} icon={icon} />
    ))}
  </div>
);

Sidebar.propTypes = {
  tab: propTypes.string.isRequired,
  setTab: propTypes.func.isRequired,
};

export default Sidebar;
