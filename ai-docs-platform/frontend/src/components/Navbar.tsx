import React from 'react';
import { Layout, Shield, FileText, Settings, User } from 'lucide-react';

const Navbar = () => {
  return (
    <nav className="fixed top-0 left-0 right-0 h-16 glass z-50 flex items-center justify-between px-6 border-b border-white/5">
      <div className="flex items-center gap-2">
        <div className="w-8 h-8 bg-primary rounded-lg flex items-center justify-center">
          <Layout className="w-5 h-5 text-white" />
        </div>
        <span className="text-xl font-bold tracking-tight text-white">AI Docs Assistant</span>
      </div>
      
      <div className="hidden md:flex items-center gap-8 text-sm font-medium text-muted-foreground">
        <a href="#" className="text-white hover:text-primary transition-colors">Projects</a>
        <a href="#" className="hover:text-primary transition-colors">Templates</a>
        <a href="#" className="hover:text-primary transition-colors">Agents</a>
        <a href="#" className="hover:text-primary transition-colors">Settings</a>
      </div>

      <div className="flex items-center gap-4">
        <div className="flex items-center gap-2 px-3 py-1.5 rounded-full bg-secondary/50 border border-white/5">
          <Shield className="w-4 h-4 text-emerald-500" />
          <span className="text-xs font-semibold text-emerald-500">Local AI Active</span>
        </div>
        <div className="w-10 h-10 rounded-full bg-secondary border border-white/10 flex items-center justify-center overflow-hidden">
          <User className="w-5 h-5 text-muted-foreground" />
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
