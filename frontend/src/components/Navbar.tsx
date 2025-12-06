import { Menu, Ticket } from "lucide-react";

const Navbar = () => {
  return (
    <nav className="bg-white shadow-sm fixed w-full z-50 top-0 border-b border-gray-100">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16 items-center">
          <div className="flex-shrink-0 flex items-center gap-2 cursor-pointer">
            <Ticket className="h-8 w-8 text-indigo-600" />
            <span className="text-2xl font-bold bg-gradient-to-r from-indigo-600 to-purple-600 bg-clip-text text-transparent">
              Nexus Events
            </span>
          </div>

          <div className="hidden md:flex items-center space-x-8">
            <a href="#" className="text-gray-600 hover:text-indigo-600 font-medium transition-colors">
              Explorar
            </a>
            <a href="#" className="text-gray-600 hover:text-indigo-600 font-medium transition-colors">
              Categorías
            </a>
            <div className="h-6 w-px bg-gray-200"></div>
            <a href="#" className="text-gray-600 hover:text-indigo-600 font-medium transition-colors">
              Iniciar Sesión
            </a>
            <button className="bg-indigo-600 text-white px-5 py-2 rounded-full font-medium hover:bg-indigo-700 transition-all shadow-md hover:shadow-lg">
              Registrarse
            </button>
          </div>

          <div className="md:hidden flex items-center">
            <button className="text-gray-500 hover:text-gray-700 focus:outline-none">
              <Menu className="h-6 w-6" />
            </button>
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;