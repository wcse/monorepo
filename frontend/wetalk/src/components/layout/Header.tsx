import Account from './Account';

export const Header = () => {
  return (
    <div className="h-16 text-white main-header">
      <div className="container mx-auto flex align-center items-center justify-between h-full">
        <div className="text-center text-3xl font-extrabold leading-none tracking-tight">
          <span className="bg-clip-text text-transparent bg-gradient-to-r from-teal-400 to-yellow-500">
            LOGO
          </span>
        </div>
        {/* <div className="w-1/3">
          <input
            type="text"
            className="appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-blue-500"
            placeholder="Search here. But it's not working"
          />
        </div> */}

        <Account />
      </div>
    </div>
  );
};

export default Header;
