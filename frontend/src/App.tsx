import { useState } from "react";
import Navbar from "./components/Navbar";
import Hero from "./sections/Hero";
import EventCard from "./components/EventCard";
import LoginForm from "./components/LoginForm";
import RegisterForm from "./components/RegisterForm"; // <--- Lo crearemos ahora
import { events } from "./data/mockData";
import { ArrowRight } from "lucide-react";
import type { EventItem } from "./types"; 
import { useAuth } from "./context/AuthContext";

function App() {
  const { isAuthenticated } = useAuth();
  const [authView, setAuthView] = useState<"login" | "register">("login");

  // USUARIO NO AUTENTICADO
  if (!isAuthenticated) {
    return authView === "login" ? (
      <LoginForm onSwitchRegister={() => setAuthView("register")} />
    ) : (
      <RegisterForm onSwitchLogin={() => setAuthView("login")} />
    );
  }

  // USUARIO AUTENTICADO
  return (
    <div className="bg-gray-50 min-h-screen font-sans">
      <Navbar />
      
      <main>
        <Hero />

        <section className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16">

          <div className="flex justify-between items-end mb-10">
            <div>
              <h2 className="text-3xl font-bold text-gray-900 mb-2">Próximos Eventos</h2>
              <p className="text-gray-600 text-lg">Explora lo que está pasando cerca de ti.</p>
            </div>
            <button className="hidden md:flex items-center text-indigo-600 hover:text-indigo-700 font-semibold transition group">
              Ver todos <ArrowRight className="ml-2 h-5 w-5 group-hover:translate-x-1 transition-transform"/>
            </button>
          </div>

          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6 lg:gap-8">
            {events.map((event: EventItem) => (
              <EventCard key={event.id} event={event} />
            ))}
          </div>

           <div className="mt-12 flex justify-center md:hidden">
               <button className="w-full sm:w-auto flex items-center justify-center bg-white border border-gray-300 px-6 py-3 rounded-full text-gray-700 font-semibold hover:bg-gray-50 transition shadow-sm">
                  Ver todos los eventos <ArrowRight className="ml-2 h-5 w-5"/>
               </button>
           </div>
        </section>
      </main>

      <footer className="bg-gray-900 text-white py-12 text-center">
          <p className="text-gray-400">© 2025 Nexus Events. Proyecto de Ingeniería.</p>
      </footer>
    </div>
  );
}

export default App;