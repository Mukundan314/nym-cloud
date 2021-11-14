import { faFolder, faClock, faStar } from "@fortawesome/free-solid-svg-icons";

export const tabs = [
  {
    name: "All Files",
    id: "all-files",
    icon: faFolder,
  },
  {
    name: "Recents",
    id: "recents",
    icon: faClock,
  },
  {
    name: "Starred",
    id: "starred",
    icon: faStar,
  },
];

export default { tabs };
